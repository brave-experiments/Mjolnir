package builder

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

	result, err := Build(staticVariablesMap)
	assert.Nil(t, err)
	assert.Equal(
		t,
		html.UnescapeString("\npackage builder\n\nvar (\nkey1 = &#96;var1&#96;\n)\n"),
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
	file, err := os.Create("dupa.go")
	assert.Nil(t, err)
	_, err = file.Write([]byte(result))
	assert.Nil(t, err)
	assert.Equal(t, "", result)

	err = os.Chdir(currentDir)
	assert.Nil(t, err)
}
