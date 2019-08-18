package terra

import (
	"fmt"
	"github.com/johandry/terranova"
	"github.com/stretchr/testify/assert"
	"github.com/terraform-providers/terraform-provider-aws/aws"
	"os"
	"testing"
)

const (
	DummyRecipeBodyFail = `variable "count"    { default = 2 }
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

func TestClient_ApplyFailure(t *testing.T) {
	dummyRecipeFileName := "dummyRecipe.tf"
	file, err := os.Create(dummyRecipeFileName)
	assert.Nil(t, err)

	num, err := file.Write([]byte(DummyRecipeBodyFail))
	assert.Nil(t, err)
	assert.Equal(t, len(DummyRecipeBodyFail), num)

	client := createTestedDefaultClient(t)
	assert.IsType(t, Client{}, client)

	variables := make(map[string]interface{})
	// Here lays type'o to stop apply on parsing tf code
	variables["key_name1"] = "dummyKey"

	recipe := File{
		Location:  dummyRecipeFileName,
		Variables: variables,
	}

	err = client.Apply(recipe, false)
	assert.Error(t, err)
	assert.Equal(
		t,
		"1 error occurred:\n\t* provider.aws: 1:3: unknown variable accessed: var.region in:\n\n${var.region}\n\n",
		err.Error(),
	)

	removeStateFileAndRestore(t)
}

func TestClient_DefaultClientCreateStateFile(t *testing.T) {
	client := createTestedDefaultClient(t)
	assert.IsType(t, Client{}, client)
	assert.FileExists(t, client.state.Location)

	err := client.state.ReadFile()
	assert.Nil(t, err)

	removeStateFileAndRestore(t)
}

func TestDefaultProvider(t *testing.T) {
	keyToTest := "dummy"
	key, provider := DefaultProvider(keyToTest)
	assert.Equal(t, keyToTest, key)
	assert.IsType(t, aws.Provider(), provider)
}

func TestClient_DumpVariables_NilVariables(t *testing.T) {
	client := createTestedDefaultClient(t)
	variables, err := client.DumpVariables()
	assert.Nil(t, err)
	assert.Empty(t, variables)

	removeStateFileAndRestore(t)
}

func TestClient_DumpVariables(t *testing.T) {
	StateFileName = "dummy.tfstate"
	stateFile, err := DefaultStateFile()
	assert.Nil(t, err)

	vars := make(map[string]interface{})
	vars["dummyKey"] = "dummyVar"
	vars["dummyKey1"] = "dummyVar1"

	platform := &terranova.Platform{
		Vars: vars,
	}

	client := Client{
		platform: platform,
		state:    stateFile,
	}

	variables, err := client.DumpVariables()
	assert.Empty(t, err)
	assert.Equal(t, vars, variables)

	removeStateFileAndRestore(t)
}

func TestClient_PreparePlatformFailure_RecipeDoesNotExist(t *testing.T) {
	fileName := "dummy.tf"
	client := Client{}
	file := File{
		Location: fileName,
	}
	err := client.PreparePlatform(file)
	assert.Error(t, err)
	assert.IsType(t, &os.PathError{}, err)
	assert.Equal(
		t,
		err.Error(),
		fmt.Sprintf("open %s: no such file or directory", fileName),
	)
}

func TestClient_PreparePlatformFailure_PlatformIsNotInitialized(t *testing.T) {
	fileName := "dummyRecipe.tf"
	fileBody := "dummy file body"
	PrepareDummyFile(t, fileName, fileBody)
	client := Client{}
	file := File{
		Location: fileName,
	}
	err := client.PreparePlatform(file)
	assert.Error(t, err)
	assert.IsType(t, ClientError{}, err)
	assert.Equal(t, "Platform is not initialized", err.Error())
	RemoveDummyFile(t, fileName)
}

func TestClient_PreparePlatformWithVariables(t *testing.T) {
	StateFileName = "dummy.tfstate"
	stateFile, err := DefaultStateFile()
	assert.Nil(t, err)

	fileName := "dummyRecipe.tf"
	fileBody := "dummy file body"
	PrepareDummyFile(t, fileName, fileBody)

	vars := make(map[string]interface{})
	vars["dummyKey"] = "dummyVar"
	vars["dummyKey1"] = "dummyVar1"

	newVars := make(map[string]interface{})
	vars["dummyKey"] = []string{"some", "values"}
	vars["dummyKey1"] = map[string]string{"dummySubKey1": "newValue"}

	// Join two maps
	joinedVars := newVars

	for key, value := range vars {
		joinedVars[key] = value
	}

	platform := &terranova.Platform{
		Vars: vars,
	}

	client := Client{
		platform: platform,
		state:    stateFile,
	}

	file := File{
		Location:  fileName,
		Variables: newVars,
	}

	err = client.PreparePlatform(file)
	assert.Nil(t, err)
	assert.Equal(t, client.platform.Code, fileBody)

	dumpedVariables, err := client.DumpVariables()
	assert.Nil(t, err)
	assert.Equal(t, dumpedVariables, joinedVars)

	RemoveDummyFile(t, fileName)
	removeStateFileAndRestore(t)
}

func TestClient_WriteStateToFileFailure(t *testing.T) {
	client := Client{
		platform: &terranova.Platform{},
	}
	err := client.WriteStateToFile()
	assert.Error(t, err)
	assert.IsType(t, ClientError{}, err)
	assert.Equal(t, "No state file found", err.Error())
}

func TestClient_WriteStateToFile(t *testing.T) {
	StateFileName = "dummy.tfstate"
	stateFile, err := DefaultStateFile()
	assert.Nil(t, err)

	platform := &terranova.Platform{}

	client := Client{
		platform: platform,
		state:    stateFile,
	}

	err = client.WriteStateToFile()
	assert.Nil(t, err)

	removeStateFileAndRestore(t)
}

func createTestedDefaultClient(t *testing.T) Client {
	StateFileName = "dummy.tfstate"
	client := Client{}
	err := client.DefaultClient()
	assert.Nil(t, err)

	return client
}

func removeStateFileAndRestore(t *testing.T) {
	RemoveDummyFile(t, StateFileName)
	StateFileName = DefaulStateFileName
}
