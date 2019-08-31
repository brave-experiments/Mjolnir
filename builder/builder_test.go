package builder

import (
	"github.com/brave-experiments/apollo-devops/src/github.com/stretchr/testify/assert"
	"github.com/brave-experiments/apollo-devops/terra"
	"testing"
)

func TestBuild(t *testing.T) {
	staticVariablesMap := map[string]interface{}{
		"StaticVariables": map[string]string{
			"key1": "var1",
		},
	}

	result, err := Build("builder", staticVariablesMap)
	assert.Nil(t, err)
	assert.Equal(t, "\npackage builder\n\nvar (\nkey1 = var1\n)\n", result)
}

func TestBuild_BuildRecipe(t *testing.T) {
	recipes := terra.DefaultRecipes
	staticVariables := make(map[string]string)

	for key, recipe := range recipes {
		err := recipe.ParseBody()
		assert.Nil(t, err)
		staticVariables[key] = recipe.Body
	}

	staticVariablesMap := map[string]interface{}{
		"StaticVariables": staticVariables,
	}

	result, err := Build("terra", staticVariablesMap)
	assert.Nil(t, err)
	assert.Nil(t, result)
}
