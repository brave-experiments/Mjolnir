package terra

import (
	"github.com/stretchr/testify/assert"
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

func TestRecipesError_Error(t *testing.T) {
	errorMsg := "Dummy Error"
	err := RecipesError{errorMsg}
	assert.IsType(t, RecipesError{}, err)
	assert.Equal(t, err.Error(), errorMsg)
}
