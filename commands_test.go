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
	YamlV1Fixture = `version: 0.1
resourceType: variables
variables: 
  simpleKey: variable
`
	ExpectedEnvKey = "APP_DEFAULT_KEY"
)

func TestApplyCmdFactory(t *testing.T) {
	command, err := ApplyCmdFactory()
	assert.Nil(t, err)
	assert.IsType(t, ApplyCmd{}, command)
	testThatCommandHasWholeInterface(t, command)
}

func TestApplyCmd_RunInvalid(t *testing.T) {
	command, err := ApplyCmdFactory()
	expectedEnvKey := ExpectedEnvKey
	expectedOldEnvValue := os.Getenv(expectedEnvKey)
	assert.Nil(t, err)
	assert.IsType(t, ApplyCmd{}, command)

	// ApplyCmd has no arguments
	exitCode := command.Run([]string{})
	assert.Equal(t, ExitCodeInvalidNumOfArgs, exitCode)

	// ApplyCmd has no Recipes
	invalidCmd := ApplyCmd{}
	dummyArgs := []string{"dummy", "dummy.yml"}
	exitCode = invalidCmd.Run(dummyArgs)
	assert.Equal(t, ExitCodeInvalidSetup, exitCode)

	// ApplyCmd has no Elements in Recipes
	recipes := terra.Recipes{}
	invalidCmd = ApplyCmd{
		Recipes: recipes,
	}
	exitCode = invalidCmd.Run(dummyArgs)
	assert.Equal(t, ExitCodeInvalidSetup, exitCode)

	// ApplyCmd has no matching key
	exitCode = command.Run(dummyArgs)
	assert.Equal(t, ExitCodeInvalidArgument, exitCode)

	// Should return error because there is no .yml variable file
	keyName := "dummy"
	filePath := "dummy.tf"
	recipes = GetMockedRecipes(
		t,
		keyName,
		filePath,
		DummyFileTfBody,
		map[string]string{},
	)
	invalidCmd = ApplyCmd{
		Recipes: recipes,
	}
	assert.IsType(t, ApplyCmd{}, invalidCmd)
	oldRecipes := terra.DefaultRecipes
	terra.DefaultRecipes = recipes.Elements
	exitCode = invalidCmd.Run(dummyArgs)
	assert.Equal(t, ExitCodeYamlBindingError, exitCode)
	terra.DefaultRecipes = oldRecipes

	// Since it is not mocked we want to end our testing process here
	yamlFileName := dummyArgs[1]
	PrepareDummyFile(t, yamlFileName, YamlV1Fixture)
	envRecipesMapping := map[string]string{
		"simpleKey": expectedEnvKey,
	}
	recipes = GetMockedRecipes(
		t,
		keyName,
		filePath,
		DummyFileTfBody,
		envRecipesMapping,
	)
	oldRecipes = terra.DefaultRecipes
	terra.DefaultRecipes = recipes.Elements
	command = ApplyCmd{
		Recipes: recipes,
	}
	exitCode = command.Run(dummyArgs)
	assert.Equal(t, ExitCodeTerraformError, exitCode)
	assert.Equal(t, expectedOldEnvValue, os.Getenv(expectedEnvKey))
	RemoveDummyFile(t, filePath)
	RemoveDummyFile(t, yamlFileName)
	terra.DefaultRecipes = oldRecipes
}

func TestApplyCmd_Run(t *testing.T) {
	// If body of file is empty it wont fail with errors
	keyName := "dummy"
	filePath := "dummy.tf"
	schemaFilePath := "dummy.yml"
	yamlFileSchema := terra.SchemaV1
	PrepareDummyFile(t, schemaFilePath, yamlFileSchema)
	recipes := GetMockedRecipes(t, keyName, filePath, "", map[string]string{})
	command := ApplyCmd{
		Recipes: recipes,
	}
	assert.IsType(t, ApplyCmd{}, command)
	exitCode := command.Run([]string{keyName, schemaFilePath})
	assert.Equal(t, ExitCodeTerraformError, exitCode)
	RemoveDummyFile(t, filePath)
	RemoveDummyFile(t, schemaFilePath)
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

func GetMockedRecipes(
	t *testing.T,
	keyName string,
	fileName string,
	fileBody string,
	envVariablesRollback map[string]string,
) (recipes terra.Recipes) {
	recipes = terra.Recipes{}
	PrepareDummyFile(t, fileName, fileBody)
	err := recipes.AddRecipe(
		keyName,
		terra.CombinedRecipe{
			File: terra.File{
				Location: fileName,
				EnvVariablesRollBack: envVariablesRollback,
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
