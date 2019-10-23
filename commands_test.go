package main

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/brave-experiments/Mjolnir/terra"
	"github.com/mitchellh/cli"
	"github.com/stretchr/testify/assert"
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

func TestNodeSshCmdFactory(t *testing.T) {
	command, err := NodeSshCmdFactory()
	assert.Nil(t, err)
	assert.IsType(t, NodeSshCmd{}, command)
	testThatCommandHasWholeInterface(t, command)
}

func TestNodeInfoSshCmdFactory(t *testing.T) {
	command, err := NodeInfoSshCmdFactory()
	assert.Nil(t, err)
	assert.IsType(t, NodeInfoSshCmd{}, command)
	testThatCommandHasWholeInterface(t, command)
}

func TestSshCmdFactory(t *testing.T) {
	command, err := SshCmdFactory()
	assert.Nil(t, err)
	assert.IsType(t, SshCmd{}, command)
	testThatCommandHasWholeInterface(t, command)
}

func TestApplyCmdFactory(t *testing.T) {
	command, err := ApplyCmdFactory()
	assert.Nil(t, err)
	assert.IsType(t, ApplyCmd{}, command)
	testThatCommandHasWholeInterface(t, command)
}

func TestDestroyCmdFactory(t *testing.T) {
	command, err := DestroyCmdFactory()
	assert.Nil(t, err)
	assert.IsType(t, DestroyCmd{}, command)
	testThatCommandHasWholeInterface(t, command)
}

func TestGethCmdFactory(t *testing.T) {
	command, err := GethCmdFactory()
	assert.Nil(t, err)
	assert.IsType(t, GethCmd{}, command)
	testThatCommandHasWholeInterface(t, command)
}

func TestGethCmd_Run(t *testing.T) {
	terra.TempDirPathLocation = ".dummyMjolnirGeth"
	dummyFileName := "output.log"
	deployName := "dummyDeployName"
	dummyDeployName := terra.TempDirPathLocation + "/" + deployName
	err := os.MkdirAll(dummyDeployName, 0777)
	assert.Nil(t, err)
	PrepareDummyFile(t, dummyDeployName+"/"+dummyFileName, terra.OutputAsAStringWithoutHeaderFixture)

	command, err := GethCmdFactory()
	assert.Nil(t, err)
	assert.IsType(t, GethCmd{}, command)

	// Should throw an sshDialError when no node number present
	exitCode := command.Run([]string{})
	assert.Equal(t, ExitCodeInvalidNumOfArgs, exitCode)

	sshKeyFileLocator := dummyDeployName + "/id_rsa"
	PrepareDummyFile(t, sshKeyFileLocator, "dummyBody")
	runArgs := []string{"a"}
	exitCode = command.Run(runArgs)
	assert.Equal(t, ExitCodeInvalidArgument, exitCode)

	runArgs = []string{"1"}
	exitCode = command.Run(runArgs)
	assert.Equal(t, ExitCodeSshDialError, exitCode)

	err = os.RemoveAll(terra.TempDirPathLocation)
	assert.Nil(t, err)
	terra.TempDirPathLocation = terra.TempDirPath
}

func TestNodeSshCmd_RunInvalid(t *testing.T) {
	terra.TempDirPathLocation = ".dummyMjolnirNode"
	dummyFileName := "output.log"
	deployName := "dummyDeployName"
	dummyDeployName := terra.TempDirPathLocation + "/" + deployName
	err := os.MkdirAll(dummyDeployName, 0777)
	assert.Nil(t, err)
	PrepareDummyFile(t, dummyDeployName+"/"+dummyFileName, terra.OutputAsAStringWithoutHeaderFixture)

	command, err := NodeSshCmdFactory()
	assert.Nil(t, err)
	assert.IsType(t, NodeSshCmd{}, command)

	// Should throw an sshDialError when no node number present
	exitCode := command.Run([]string{})
	assert.Equal(t, ExitCodeInvalidNumOfArgs, exitCode)

	sshKeyFileLocator := dummyDeployName + "/id_rsa"
	PrepareDummyFile(t, sshKeyFileLocator, "dummyBody")
	runArgs := []string{"a"}
	exitCode = command.Run(runArgs)
	assert.Equal(t, ExitCodeInvalidArgument, exitCode)

	runArgs = []string{"1"}
	exitCode = command.Run(runArgs)
	assert.Equal(t, ExitCodeSshDialError, exitCode)

	err = os.RemoveAll(terra.TempDirPathLocation)
	assert.Nil(t, err)
	terra.TempDirPathLocation = terra.TempDirPath
}

func TestNodeInfoSshCmd_RunInvalid(t *testing.T) {
	terra.TempDirPathLocation = ".dummyMjolnirNodeInfo"
	dummyFileName := "output.log"
	deployName := "dummyDeployName"
	dummyDeployName := terra.TempDirPathLocation + "/" + deployName
	err := os.MkdirAll(dummyDeployName, 0777)
	assert.Nil(t, err)
	PrepareDummyFile(t, dummyDeployName+"/"+dummyFileName, terra.OutputAsAStringWithoutHeaderFixture)

	command, err := NodeInfoSshCmdFactory()
	assert.Nil(t, err)
	assert.IsType(t, NodeInfoSshCmd{}, command)

	exitCode := command.Run([]string{})
	assert.Equal(t, ExitCodeSshDialError, exitCode)

	err = os.RemoveAll(terra.TempDirPathLocation)
	assert.Nil(t, err)
	terra.TempDirPathLocation = terra.TempDirPath
}

func TestSshCmd_RunInvalid(t *testing.T) {
	terra.TempDirPathLocation = ".dummyMjolnir"
	dummyFileName := "output.log"
	deployName := "dummyDeployName"
	dummyDeployName := terra.TempDirPathLocation + "/" + deployName
	err := os.MkdirAll(dummyDeployName, 0777)
	assert.Nil(t, err)
	PrepareDummyFile(t, dummyDeployName+"/"+dummyFileName, terra.OutputAsAStringWithoutHeaderFixture)

	command, err := SshCmdFactory()
	assert.Nil(t, err)
	assert.IsType(t, SshCmd{}, command)

	// Should throw an sshDialError when no key present
	exitCode := command.Run([]string{"-i", "dummyFile"})
	assert.Equal(t, ExitCodeSshDialError, exitCode)

	sshKeyFileLocator := dummyDeployName + "/id_rsa"
	PrepareDummyFile(t, sshKeyFileLocator, "dummyBody")
	runArgs := []string{"-i", "dummyFile"}
	exitCode = command.Run(runArgs)
	assert.Equal(t, ExitCodeSshDialError, exitCode)

	err = os.RemoveAll(terra.TempDirPathLocation)
	assert.Nil(t, err)
	terra.TempDirPathLocation = terra.TempDirPath
}

func TestApplyCmd_RunInvalid(t *testing.T) {
	terra.TempDirPathLocation = ".mjolnirApplyTemp"
	err := os.RemoveAll(terra.TempDirPathLocation)
	assert.Nil(t, err)
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
	err = os.RemoveAll(terra.TempDirPathLocation)
	assert.Nil(t, err)
	terra.TempDirPathLocation = terra.TempDirPath
}

func TestDestroyCmd_RunInvalid(t *testing.T) {
	terra.TempDirPathLocation = ".mjolnirTestDir"
	err := os.RemoveAll(terra.TempDirPathLocation)
	assert.Nil(t, err)

	command, err := DestroyCmdFactory()
	expectedEnvKey := ExpectedEnvKey
	expectedOldEnvValue := os.Getenv(expectedEnvKey)
	assert.Nil(t, err)
	assert.IsType(t, DestroyCmd{}, command)

	// DestroyCmd has no arguments
	exitCode := command.Run([]string{})
	assert.Equal(t, ExitCodeInvalidNumOfArgs, exitCode)
	_, err = os.Stat(terra.TempDirPathLocation)
	assert.True(t, os.IsNotExist(err))

	// DestroyCmd has no Recipes
	invalidCmd := DestroyCmd{}
	dummyArgs := []string{"dummy.yml"}
	exitCode = invalidCmd.Run(dummyArgs)
	assert.Equal(t, ExitCodeYamlBindingError, exitCode)
	_, err = os.Stat(terra.TempDirPathLocation)
	assert.True(t, os.IsNotExist(err))

	// DestroyCmd has no Elements in Recipes
	recipes := terra.Recipes{}
	invalidDestroyCmd := DestroyCmd{
		ApplyCmd{
			Recipes: recipes,
		},
	}
	exitCode = invalidDestroyCmd.Run(dummyArgs)
	assert.Equal(t, ExitCodeYamlBindingError, exitCode)
	_, err = os.Stat(terra.TempDirPathLocation)
	assert.True(t, os.IsNotExist(err))

	// DestroyCmd has no matching key
	exitCode = command.Run(dummyArgs)
	assert.Equal(t, ExitCodeYamlBindingError, exitCode)
	_, err = os.Stat(terra.TempDirPathLocation)
	assert.True(t, os.IsNotExist(err))

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
	invalidDestroyCmd = DestroyCmd{
		ApplyCmd{
			Recipes: recipes,
		},
	}
	assert.IsType(t, DestroyCmd{}, invalidDestroyCmd)
	oldRecipes := terra.DefaultRecipes
	terra.DefaultRecipes = recipes.Elements
	exitCode = invalidDestroyCmd.Run(dummyArgs)
	assert.Equal(t, ExitCodeYamlBindingError, exitCode)
	terra.DefaultRecipes = oldRecipes

	// Since it is not mocked we want to end our testing process here
	err = os.MkdirAll(terra.TempDirPathLocation, 0777)
	assert.Nil(t, err)
	yamlFileName := dummyArgs[0]
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
	terra.DestroyDefaultRecipeVar = "dummy"
	command = DestroyCmd{
		ApplyCmd{Recipes: recipes},
	}
	exitCode = command.Run(dummyArgs)
	assert.Equal(t, ExitCodeSuccess, exitCode)
	assert.Equal(t, expectedOldEnvValue, os.Getenv(expectedEnvKey))
	terra.DestroyDefaultRecipeVar = terra.DestroyDefaultRecipe
	_, err = os.Stat(terra.TempDirPathLocation)
	assert.True(t, os.IsNotExist(err))
	RemoveDummyFile(t, filePath)
	RemoveDummyFile(t, yamlFileName)
	terra.DefaultRecipes = oldRecipes
	err = os.RemoveAll(terra.TempDirPathLocation)
	assert.Nil(t, err)
	terra.TempDirPathLocation = terra.TempDirPath
}

func TestApplyCmd_Run(t *testing.T) {
	terra.TempDirPathLocation = ".mjolnirApplyEnd"
	err := os.RemoveAll(terra.TempDirPathLocation)
	assert.Nil(t, err)
	// We want to end with ExitCodeTerraformError without actual e2e calls
	keyName := "dummy"
	filePath := "dummy.tf"
	schemaFilePath := "dummy.yml"
	yamlFileSchema := terra.SchemaV02
	PrepareDummyFile(t, schemaFilePath, yamlFileSchema)
	recipes := GetMockedRecipes(t, keyName, filePath, "", map[string]string{})
	command := ApplyCmd{
		Recipes: recipes,
	}
	assert.IsType(t, ApplyCmd{}, command)
	exitCode := command.Run([]string{keyName, schemaFilePath})
	assert.Equal(t, ExitCodeSuccess, exitCode)
	RemoveDummyFile(t, filePath)
	RemoveDummyFile(t, schemaFilePath)
	err = os.RemoveAll(terra.TempDirPathLocation)
	assert.Nil(t, err)
	terra.TempDirPathLocation = terra.TempDirPath
}

func TestDestroyCmd_Run(t *testing.T) {
	terra.TempDirPathLocation = ".mjolnirTestTemp"
	err := os.RemoveAll(terra.TempDirPathLocation)
	assert.Nil(t, err)
	keyName := "dummy"
	filePath := "dummy.tf"
	schemaFilePath := "dummy.yml"
	yamlFileSchema := terra.SchemaV02
	PrepareDummyFile(t, schemaFilePath, yamlFileSchema)
	recipes := GetMockedRecipes(t, keyName, filePath, "", map[string]string{})
	commandDestroy := DestroyCmd{
		ApplyCmd{
			Recipes: recipes,
		},
	}
	assert.IsType(t, DestroyCmd{}, commandDestroy)
	exitCode := commandDestroy.Run([]string{keyName, schemaFilePath})
	assert.Equal(t, ExitCodeYamlBindingError, exitCode)
	_, err = os.Stat(terra.TempDirPathLocation)
	assert.True(t, os.IsNotExist(err))
	RemoveDummyFile(t, filePath)
	RemoveDummyFile(t, schemaFilePath)
	terra.TempDirPathLocation = terra.TempDirPath
}

func TestApplyCmd_Help(t *testing.T) {
	command, err := ApplyCmdFactory()
	assert.Nil(t, err)
	assert.IsType(t, ApplyCmd{}, command)
	helpMsg := command.Help()
	assert.Greater(t, len(helpMsg), 0)
}

func TestDestroyCmd_Help(t *testing.T) {
	command, err := DestroyCmdFactory()
	assert.Nil(t, err)
	assert.IsType(t, DestroyCmd{}, command)
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
				Location:             fileName,
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
