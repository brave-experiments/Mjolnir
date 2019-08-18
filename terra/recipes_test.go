package terra

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"testing"
)

func TestRecipes_CreateWithDefaults(t *testing.T) {
	recipes := Recipes{}
	recipes.CreateWithDefaults()
	assert.Greater(t, len(recipes.Elements), 0)
}

func TestRecipes_AddRecipe(t *testing.T) {
	keyName := "dummy"
	dummyRecipe := File{Location: keyName}
	recipes := Recipes{}
	err := recipes.AddRecipe(keyName, dummyRecipe)
	assert.Nil(t, err)
	assert.Equal(t, recipes.Elements[keyName], dummyRecipe)
}

func TestRecipes_AddRecipeToDefaults(t *testing.T) {
	keyName := "dummy"
	dummyRecipe := File{Location: keyName}
	recipes := Recipes{}
	recipes.CreateWithDefaults()
	err := recipes.AddRecipe(keyName, dummyRecipe)
	assert.Nil(t, err)
	assert.Equal(t, recipes.Elements[keyName], dummyRecipe)
}

func TestRecipes_AddRecipeFailure(t *testing.T) {
	keyName := "dummy"
	dummyRecipe := File{Location: keyName}
	recipes := Recipes{}
	err := recipes.AddRecipe(keyName, dummyRecipe)
	assert.Nil(t, err)
	assert.Equal(t, recipes.Elements[keyName], dummyRecipe)

	err = recipes.AddRecipe(keyName, dummyRecipe)
	// Test that list returned error caused by existing key
	assert.IsType(t, RecipesError{}, err)
	assert.Equal(t, err.Error(), fmt.Sprintf("%s  already exists in recipes list", keyName))
	// Test that list is immutable to error
	assert.Equal(t, recipes.Elements[keyName], dummyRecipe)
}

func TestRecipesError_Error(t *testing.T) {
	errorMsg := "Dummy Error"
	err := RecipesError{errorMsg}
	assert.IsType(t, RecipesError{}, err)
	assert.Equal(t, err.Error(), errorMsg)
}

func TestClientError_Error(t *testing.T) {
	errorMsg := "Dummy Msg"
	err := RecipesError{errorMsg}
	assert.IsType(t, RecipesError{}, err)
	assert.Equal(t, err.Error(), errorMsg)
}

func TestFile_ReadFileFailure(t *testing.T) {
	file := File{}
	err := file.ReadFile()
	assert.Error(t, err)
	assert.Equal(t, err.Error(), "open : no such file or directory")
}

func TestFile_ReadFile(t *testing.T) {
	fileBody := "dummy content"
	fileName := "dummyFileUniqueName.tf"
	PrepareDummyFile(t, fileName, fileBody)

	file := File{
		Location: fileName,
	}

	err := file.ReadFile()
	assert.Nil(t, err)

	assert.Equal(t, string(fileBody), file.Body)

	RemoveDummyFile(t, fileName)
}

func TestFile_ReadFileWithVariables(t *testing.T) {
	// This test is done to concrete variables design
	variables := map[string]interface{}{
		"dummyKey": "dummyVal",
	}

	file := File{
		Variables: variables,
	}

	assert.IsType(t, file, File{})
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
