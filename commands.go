package main

import (
    "fmt"
    "github.com/brave-experiments/apollo-devops/terra"
    "github.com/mitchellh/cli"
)

const (
    ExitCodeSuccess = 0
    ExitCodeInvalidNumOfArgs = 1
    ExitCodeInvalidSetup = 2
    ExitCodeInvalidArgument = 3
    ExitCodeTerraformError = 4
)

var (
    RegisteredCommands = map[string]cli.CommandFactory{
        "apply": ApplyCmdFactory,
    }
)

type ApplyCmd struct {
    cli.Command
    Recipes terra.Recipes
}

func ApplyCmdFactory() (command cli.Command, err error) {
    recipes := terra.Recipes{}
    recipes.CreateWithDefaults()

    command = ApplyCmd{
        Recipes: recipes,
    }

    return command, err
}

func (applyCmd ApplyCmd) Run(args []string) (exitCode int) {
    expectedMinimumArguments := 1

    if len(args) < expectedMinimumArguments {
        fmt.Printf(
            "Not enough arguments, expected more than %v \n",
            expectedMinimumArguments,
        )

        return ExitCodeInvalidNumOfArgs
    }

    if nil == &applyCmd.Recipes || len(applyCmd.Recipes.Elements) < 1 {
        fmt.Println("Invalid recipes configuration")

        return ExitCodeInvalidSetup
    }

    recipes := applyCmd.Recipes.Elements
    recipeKey := args[0]
    recipe, contains := recipes[recipeKey]

    if false == contains {
        fmt.Printf(
            "Recipe %s not found within recipes, available are: \n",
            recipeKey,
        )
        applyCmd.printRecipesKeys()

        return ExitCodeInvalidArgument
    }

    fmt.Printf(
        "Executing %s \n",
        recipeKey,
    )

    err := applyCmd.executeTerra(recipe)

    if nil != err {
        fmt.Println(err)
        return ExitCodeTerraformError
    }

    return ExitCodeSuccess
}

func (applyCmd ApplyCmd) Help() (helpMessage string) {
    helpMessage = "\nThis is apply command. Usage: apollo apply [recipe]\n"

    if nil != &applyCmd.Recipes {
        helpMessage = helpMessage + "\nAvailable Recipes: \n"

        for recipeKey := range applyCmd.Recipes.Elements {
            helpMessage = helpMessage + "\n \t" + recipeKey + "\n"
        }
    }

    return helpMessage
}

func (applyCmd ApplyCmd) Synopsis() (synopsis string) {
    synopsis = "apply [recipe]"
    return synopsis
}

func (applyCmd ApplyCmd) printRecipesKeys() {
    if nil == &applyCmd.Recipes {
        return
    }

    for key := range applyCmd.Recipes.Elements {
        fmt.Println(key)
    }
}

func (applyCmd *ApplyCmd) executeTerra(recipe terra.CombinedRecipe) (err error) {
   terraClient := terra.Client{}
   err = terraClient.DefaultClient()

   if nil != err {
       return err
   }

   err = terraClient.ApplyCombined(recipe, false)

   return err
}
