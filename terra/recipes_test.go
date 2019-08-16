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
	dummyRecipe := Recipe{Location: keyName}
	recipes := Recipes{}
	err := recipes.AddRecipe(keyName, dummyRecipe)
	assert.Nil(t, err)
	assert.Equal(t, recipes.Elements[keyName], dummyRecipe)
}

func TestRecipes_AddRecipeToDefaults(t *testing.T) {
	keyName := "dummy"
	dummyRecipe := Recipe{Location: keyName}
	recipes := Recipes{}
	recipes.CreateWithDefaults()
	err := recipes.AddRecipe(keyName, dummyRecipe)
	assert.Nil(t, err)
	assert.Equal(t, recipes.Elements[keyName], dummyRecipe)
}

func TestRecipes_AddRecipeFailure(t *testing.T) {
	keyName := "dummy"
	dummyRecipe := Recipe{Location: keyName}
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

func TestRecipe_CreateFailure(t *testing.T) {
	recipe := Recipe{}
	err := recipe.Create()
	assert.Error(t, err)
	assert.Equal(t, err.Error(), "open : no such file or directory")
}

func TestRecipe_Create(t *testing.T) {
	fileBody := []byte("dummy content")
	fileName := "dummyFileUniqueName.tf"

	err := ioutil.WriteFile(fileName, fileBody, 0644)
	assert.Nil(t, err)

	recipe := Recipe{
		Location: fileName,
	}

	err = recipe.Create()
	assert.Nil(t, err)

	assert.Equal(t, string(fileBody), recipe.Body)

	err = os.Remove(fileName)
	assert.Nil(t, err)
}
