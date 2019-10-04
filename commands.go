package main

import (
    "fmt"
    "github.com/brave-experiments/apollo-devops/terra"
    "github.com/mitchellh/cli"
    "io/ioutil"
    "os"
)

const (
    ExitCodeSuccess = 0
    ExitCodeInvalidNumOfArgs = 1
    ExitCodeInvalidSetup = 2
    ExitCodeInvalidArgument = 3
    ExitCodeTerraformError = 4
    ExitCodeYamlBindingError = 5
    ExitCodeEnvUnbindingError = 6
    ExitCodeNoTempDirectory = 7
)

var (
    RegisteredCommands = map[string]cli.CommandFactory{
        "apply": ApplyCmdFactory,
        "destroy": DestroyCmdFactory,
        "ssh": SshCmdFactory,
    }
)

type ApplyCmd struct {
    cli.Command
    Recipes terra.Recipes
}

type DestroyCmd struct {
    ApplyCmd
}

type SshCmd struct {
    cli.Command
}

func SshCmdFactory() (command cli.Command, err error) {
    return SshCmd{}, nil
}

func ApplyCmdFactory() (command cli.Command, err error) {
    recipes := terra.Recipes{}
    recipes.CreateWithDefaults()

    command = ApplyCmd{
        Recipes: recipes,
    }

    return command, err
}

func DestroyCmdFactory() (command cli.Command, err error) {
    recipes := terra.Recipes{}
    recipes.CreateWithDefaults()
    command = DestroyCmd{
        ApplyCmd: ApplyCmd{
            Recipes: recipes,
        },
    }

    return command, err
}

func (sshCmd SshCmd) Run(args []string) (exitCode int) {
    //desiredBastionKey := "bastion_host_ip"
    directoryLocator := terra.TempDirPathLocation
    directoriesList, err := ioutil.ReadDir(directoryLocator)

    if nil != err || len(directoriesList) < 1 {
       fmt.Printf("%s temp directory not present \n", directoryLocator)
       return ExitCodeNoTempDirectory
    }

    return ExitCodeSuccess
}

func (destroyCmd DestroyCmd) Run(args []string) (exitCode int) {
    expectedMinimumArguments := 1

    if len(args) < expectedMinimumArguments {
        fmt.Printf(
            "Not enough arguments, expected more than %v \n write `destroy --help` for more info\n",
            expectedMinimumArguments,
        )

        return ExitCodeInvalidNumOfArgs
    }

    yamlFilePath := args[0]
    recipe, exitCode := destroyCmd.getRecipe("quorum")

    if exitCode != ExitCodeSuccess {
        fmt.Printf("Exited with code: %v, recipe not found", exitCode)
    }

    err := recipe.BindYamlWithVars(yamlFilePath)

    if nil != err {
        return ExitCodeYamlBindingError
    }

    err = destroyCmd.ApplyCmd.executeTerra(recipe, true)
    exitCode = ExitCodeSuccess

    if nil != err {
        fmt.Println(err)
        exitCode = ExitCodeTerraformError
    }

    fmt.Println("Restoring env variables.")
    err = destroyCmd.ApplyCmd.restoreEnvVariables(recipe)

    if nil != err {
        fmt.Printf("Error restoring variables: %s", err)
        exitCode = ExitCodeEnvUnbindingError
    }

    fmt.Printf( "Destroy complete! Resources in network %s destroyed.\n", recipe.Variables["network_name"])

    return exitCode
}

func (applyCmd ApplyCmd) Run(args []string) (exitCode int) {
    expectedMinimumArguments := 2

    if len(args) < expectedMinimumArguments {
        fmt.Printf(
            "Not enough arguments, expected more than %v \n write `apply --help` for more info\n",
            expectedMinimumArguments,
        )

        return ExitCodeInvalidNumOfArgs
    }

    if nil == &applyCmd.Recipes || len(applyCmd.Recipes.Elements) < 1 {
        fmt.Println("Invalid recipes configuration")

        return ExitCodeInvalidSetup
    }

    recipeKey := args[0]
    recipe, exitCode := applyCmd.getRecipe(recipeKey)

    if exitCode > ExitCodeSuccess {
        return exitCode
    }

    yamlFilePath := args[1]
    err := recipe.BindYamlWithVars(yamlFilePath)

    if nil != err {
        return ExitCodeYamlBindingError
    }

    fmt.Printf(
        "Executing %s \n",
        recipeKey,
    )

    err = applyCmd.executeTerra(recipe, false)
    exitCode = ExitCodeSuccess

    if nil != err {
        fmt.Println(err)
        exitCode = ExitCodeTerraformError
    }

    fmt.Println("Restoring env variables.")
    err = applyCmd.restoreEnvVariables(recipe)

    if nil != err {
        fmt.Printf("Error restoring variables: %s", err)
        exitCode = ExitCodeEnvUnbindingError
    }

    return exitCode
}

func (sshCmd SshCmd) Help() (helpMessage string) {
    return "Ssh into bastion. .apollo directory must be present (previous deploy must complete)"
}

func (applyCmd ApplyCmd) Help() (helpMessage string) {
    helpMessage = "\nThis is apply command. Usage: apollo apply [recipe] [yamlFilePath]\n"

    if nil != &applyCmd.Recipes {
        helpMessage = helpMessage + "\nAvailable Recipes: \n"

        for recipeKey := range applyCmd.Recipes.Elements {
            helpMessage = helpMessage + "\n \t" + recipeKey + "\n"
        }
    }

    helpMessage = helpMessage + "\n" + "Filepath must be valid yaml file with variables, like:" + "\n"
    helpMessage = helpMessage + "\n" + terra.SchemaV02 + "\n"

    return helpMessage
}

func (destroyCmd DestroyCmd) Help() (helpMessage string) {
    helpMessage = "\nThis is destroy command. Usage: apollo destroy [yamlFilePath]"
    helpMessage = helpMessage + "\nThere must be a valid terraform.tfstate file at root of execution"

    return helpMessage
}

func (applyCmd ApplyCmd) Synopsis() (synopsis string) {
    synopsis = "apply [recipe] [yamlSchemaPath]"
    return synopsis
}

func (destroyCmd DestroyCmd) Synopsis() (synopsis string) {
    synopsis = "destroy [yamlSchemaPath]"
    return synopsis
}

func (sshCmd SshCmd) Synopsis() (synopsis string) {
    synopsis = "ssh"
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

func (applyCmd *ApplyCmd) executeTerra(recipe terra.CombinedRecipe, destroy bool) (err error) {
   terraClient := terra.Client{}
   err = terraClient.DefaultClient()

   if nil != err {
       return err
   }

   err = terraClient.ApplyCombined(recipe, destroy)

   if true == destroy {
       _  = os.RemoveAll(terra.TempDirPathLocation)
   }

   return err
}

func (applyCmd *ApplyCmd) getRecipe(recipeKey string) (recipe terra.CombinedRecipe, exitCode int) {
    recipes := applyCmd.Recipes.Elements
    recipe, contains := recipes[recipeKey]

    if false == contains {
        fmt.Printf(
            "Recipe %s not found within recipes, available are: \n",
            recipeKey,
        )
        applyCmd.printRecipesKeys()

        return terra.CombinedRecipe{}, ExitCodeInvalidArgument
    }

    return recipe, ExitCodeSuccess
}

func (applyCmd *ApplyCmd) restoreEnvVariables(recipe terra.CombinedRecipe) (err error) {
    return recipe.UnbindEnvVars()
}
