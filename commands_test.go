package main

import (
    "github.com/brave-experiments/apollo-devops/terra"
    "github.com/mitchellh/cli"
    "github.com/stretchr/testify/assert"
    "io/ioutil"
    "os"
    "testing"
)

const (
    DummyFileTfBody = `variable "count"    { default = 2 }
  variable "key_name" {}
  variable "region" {}
  provider "aws" {
    region        =  "${var.region}"
  }
  resource "aws_instance" "server" {
    instance_type = "t2.micro"
    ami           = "ami-6e1a0117"
    count         = "${var.count}"
    key_name      = "${var.key_name}"
  }`
)

func TestApplyCmdFactory(t *testing.T) {
    command, err := ApplyCmdFactory()
    assert.Nil(t, err)
    assert.IsType(t, ApplyCmd{}, command)
    testThatCommandHasWholeInterface(t, command)
}

func TestApplyCmd_RunInvalid(t *testing.T) {
    command, err := ApplyCmdFactory()
    assert.Nil(t, err)
    assert.IsType(t, ApplyCmd{}, command)

    // ApplyCmd has no arguments
    exitCode := command.Run([]string{})
    assert.Equal(t, ExitCodeInvalidNumOfArgs, exitCode)

    // ApplyCmd has no Recipes
    invalidCmd := ApplyCmd{}
    exitCode = invalidCmd.Run([]string{"dummy"})
    assert.Equal(t, ExitCodeInvalidSetup, exitCode)

    // ApplyCmd has no Elements in Recipes
    recipes := terra.Recipes{}
    invalidCmd = ApplyCmd{
        Recipes: recipes,
    }
    exitCode = invalidCmd.Run([]string{"dummy"})
    assert.Equal(t, ExitCodeInvalidSetup, exitCode)

    // ApplyCmd has no matching key
    exitCode = command.Run([]string{"dummy"})
    assert.Equal(t, ExitCodeInvalidArgument, exitCode)
    recipes = terra.Recipes{}
    recipes.CreateWithDefaults()
}

func TestApplyCmd_Run(t *testing.T) {
    keyName := "dummy"
    filePath := "dummy.tf"
    recipes := GetMockedRecipes(t, keyName, filePath)
    command := ApplyCmd{
       Recipes: recipes,
    }
    assert.IsType(t, ApplyCmd{}, command)
    terra.DefaultRecipes = recipes.Elements
    exitCode := command.Run([]string{"dummy"})
    // Since it is not mocked we want to end our testing process here
    assert.Equal(t, ExitCodeSuccess, exitCode)
    RemoveDummyFile(t, filePath)
}

func TestApplyCmd_Help(t *testing.T) {
    command, err := ApplyCmdFactory()
    assert.Nil(t, err)
    assert.IsType(t, ApplyCmd{}, command)
    helpMsg := command.Help()
    assert.Greater(t, len(helpMsg), 0)
}

func testThatCommandHasWholeInterface(t *testing.T, command cli.Command) {
    helpMsg := command.Help()
    assert.Greater(t, len(helpMsg), 0)

    exitCode := command.Run([]string{})
    assert.IsType(t, 0, exitCode)

    synopsis := command.Synopsis()
    assert.Greater(t, len(synopsis), 0)
}

func GetMockedRecipes(t *testing.T, keyName string, fileName string) (recipes terra.Recipes) {
    recipes = terra.Recipes{}
    PrepareDummyFile(t, fileName, DummyFileTfBody)
    err := recipes.AddRecipe(
        keyName,
        terra.CombinedRecipe{
          File: terra.File{
              Location: fileName,
          },
          FilePaths: []string{fileName},
        },
    )
    assert.Nil(t, err)

    return recipes
}

func PrepareDummyFile(t *testing.T, fileName string, content string) {
    fileBody := []byte(content)

    err := ioutil.WriteFile(fileName, fileBody, 0644)
    assert.Nil(t, err)
}

func RemoveDummyFile(t *testing.T, fileName string) {
    err := os.Remove(fileName)
    assert.Nil(t, err)
}
