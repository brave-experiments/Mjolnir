package main

import (
    "github.com/brave-experiments/apollo-devops/terra"
    "github.com/mitchellh/cli"
)

const (
    ExitCodeSuccess = 0
    ExitCodeInvalidNumOfArgs = 1
    ExitCodeInvalidSetup = 2
    ExitCodeInvalidArgument = 3
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
    if len(args) < 1 {
        return ExitCodeInvalidNumOfArgs
    }

    if nil == &applyCmd.Recipes || len(applyCmd.Recipes.Elements) < 1 {
        return ExitCodeInvalidSetup
    }

    recipes := applyCmd.Recipes.Elements
    recipeKey := args[0]
    _, contains := recipes[recipeKey]

    if false == contains {
        return ExitCodeInvalidArgument
    }

    return ExitCodeSuccess
}

func (applyCmd ApplyCmd) Help() (helpMessage string) {
    return "This Command is not finished yet"
}

func (applyCmd ApplyCmd) Synopsis() (synopsis string) {
    synopsis = "apply [recipe] [optional]"
    return synopsis
}

