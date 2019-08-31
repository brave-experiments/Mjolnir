package main

import (
	"github.com/brave-experiments/apollo-devops/terra"
	"github.com/stretchr/testify/assert"
	"golang.org/x/net/html"
	"os"
	"testing"
)

func TestBuild(t *testing.T) {
	staticVariablesMap := map[string]interface{}{
		"StaticVariables": map[string]string{
			"key1": "var1",
		},
	}
	// It is hardcoded, because it is crucial to further workflow of app
	expectedString := "\npackage builder\n\nvar (\nStaticKey1 = `var1`\n)\n"

	result, err := Build(staticVariablesMap)
	assert.Nil(t, err)
	assert.Equal(
		t,
		html.UnescapeString(expectedString),
		result,
	)
}

func TestBuild_BuildRecipe(t *testing.T) {
	currentDir, err := os.Getwd()
	assert.Nil(t, err)

	err = os.Chdir("../")
	assert.Nil(t, err)

	recipes := terra.DefaultRecipes
	staticVariables := make(map[string]string)

	for key, recipe := range recipes {
		err := recipe.ParseBody()
		assert.Nil(t, err)
		staticVariables[key] = recipe.Body
	}

	staticVariablesMap := map[string]interface{}{
		"PackageName":     "terra",
		"StaticVariables": staticVariables,
	}

	result, err := Build(staticVariablesMap)
	assert.Nil(t, err)
	// Check that file has more than 100 chars.
	// It means that parsing process injected combined recipes
	assert.Greater(t, len(result), 100)

	err = os.Chdir(currentDir)
	assert.Nil(t, err)
}
