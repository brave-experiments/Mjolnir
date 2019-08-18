package terra

import (
	"fmt"
	"github.com/johandry/terranova"
	"github.com/stretchr/testify/assert"
	"github.com/terraform-providers/terraform-provider-aws/aws"
	"os"
	"testing"
)

func TestClient_DefaultClientCreateStateFile(t *testing.T) {
	StateFileName = "dummy.tfstate"
	client := Client{}
	err := client.DefaultClient()
	assert.Greater(t, len(client.Recipes.Elements), 0)
	assert.Nil(t, err)

	RemoveDummyFile(t, StateFileName)
	StateFileName = DefaulStateFileName
}

func TestDefaultProvider(t *testing.T) {
	keyToTest := "dummy"
	key, provider := DefaultProvider(keyToTest)
	assert.Equal(t, keyToTest, key)
	assert.IsType(t, aws.Provider(), provider)
}

func TestClient_DumpVariables_NilVariables(t *testing.T) {
	client := Client{}
	err := client.DefaultClient()
	assert.Nil(t, err)
	variables, err := client.DumpVariables()
	assert.Nil(t, err)
	assert.Empty(t, variables)
}

func TestClient_DumpVariables(t *testing.T) {
	vars := make(map[string]interface{})
	vars["dummyKey"] = "dummyVar"
	vars["dummyKey1"] = "dummyVar1"

	platform := &terranova.Platform{
		Vars: vars,
	}

	client := Client{
		platform: platform,
	}

	variables, err := client.DumpVariables()
	assert.Empty(t, err)
	assert.Equal(t, vars, variables)
}

func TestClient_RunPlatformFailure_RecipeDoesNotExist(t *testing.T) {
	fileName := "dummy.tf"
	client := Client{}
	file := File{
		Location: fileName,
	}
	err := client.RunPlatform(file)
	assert.Error(t, err)
	assert.IsType(t, &os.PathError{}, err)
	assert.Equal(
		t,
		err.Error(),
		fmt.Sprintf("open %s: no such file or directory", fileName),
	)
}

func TestClient_RunPlatformFailure_PlatformIsNotInitialized(t *testing.T) {
	fileName := "dummyRecipe.tf"
	fileBody := "dummy file body"
	PrepareDummyFile(t, fileName, fileBody)
	client := Client{}
	file := File{
		Location: fileName,
	}
	err := client.RunPlatform(file)
	assert.Error(t, err)
	assert.IsType(t, ClientError{}, err)
	assert.Equal(t, "Platform is not initialized", err.Error())
	RemoveDummyFile(t, fileName)
}

func TestClient_RunPlatformWithVariables(t *testing.T) {
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
	}

	file := File{
		Location:  fileName,
		Variables: newVars,
	}

	err := client.RunPlatform(file)
	assert.Nil(t, err)

	dumpedVariables, err := client.DumpVariables()
	assert.Nil(t, err)
	assert.Equal(t, dumpedVariables, joinedVars)

	RemoveDummyFile(t, fileName)
}

//func TestClient_RunPlatform(t *testing.T) {
//	fileName := "dummyRecipe.tf"
//	fileBody := "dummy file body"
//	PrepareDummyFile(t, fileName, fileBody)
//
//	vars := make(map[string]interface{})
//	vars["dummyKey"] = "dummyVar"
//	vars["dummyKey1"] = "dummyVar1"
//
//	platform := &terranova.Platform{
//		Vars: vars,
//	}
//
//	client := Client{
//		platform: platform,
//	}
//
//	file := File{
//		Location: fileName,
//	}
//}
