package terra

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"testing"
)

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
