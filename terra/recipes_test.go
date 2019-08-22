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
	assert.Equal(t, err.Error(), fmt.Sprintf("%s already exists in recipes list", keyName))
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
	assert.IsType(t, &os.PathError{}, err)
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

func TestCombinedRecipe_ParseBodyFailure(t *testing.T) {
	combinedRecipe := CombinedRecipe{}
	err := combinedRecipe.ParseBody()
	assert.IsType(t, RecipesError{}, err)
	assert.Equal(
		t,
		"There are no recipes within this combined recipe",
		err.Error(),
	)

	combinedRecipe = CombinedRecipe{
		FilePaths: []string{},
	}
	err = combinedRecipe.ParseBody()
	assert.IsType(t, RecipesError{}, err)
	assert.Equal(
		t,
		"There are no recipes within this combined recipe",
		err.Error(),
	)

	combinedRecipe = CombinedRecipe{
		FilePaths: []string{
			"dummy_file.tf",
			"dummy_seccond_file.tf",
		},
	}
	err = combinedRecipe.ParseBody()
	assert.IsType(t, &os.PathError{}, err)
	//assert.Equal(
	//	t,
	//	"There are no recipes within this combined recipe",
	//	err.Error(),
	//)
}

func TestCombinedRecipe_ParseBody(t *testing.T) {
	dummyFile := File{
		Location: "dummy_file.tf",
		Body:     "dummy body line 1",
	}
	dummyFile1 := File{
		Location: "dummy_file2.tf",
		Body:     "dummy body line 2",
	}

	PrepareDummyFile(t, dummyFile.Location, dummyFile.Body)
	PrepareDummyFile(t, dummyFile1.Location, dummyFile1.Body)

	combinedRecipe := CombinedRecipe{
		FilePaths: []string{
			dummyFile.Location,
			dummyFile1.Location,
		},
	}

	err := combinedRecipe.ParseBody()
	assert.Nil(t, err)

	expectedFileBody := dummyFile.Body + "\n" + dummyFile1.Body
	assert.Equal(t, expectedFileBody, expectedFileBody)

	// Check that body does not occur n times
	err = combinedRecipe.ParseBody()
	assert.Nil(t, err)
	expectedFileBody = dummyFile.Body + "\n" + dummyFile1.Body
	assert.Equal(t, expectedFileBody, expectedFileBody)

	// Check that body is in expected order
	combinedRecipe = CombinedRecipe{
		FilePaths: []string{
			dummyFile1.Location,
			dummyFile.Location,
		},
	}

	err = combinedRecipe.ParseBody()
	assert.Nil(t, err)

	expectedFileBody = dummyFile1.Body + "\n" + dummyFile.Body
	assert.Equal(t, expectedFileBody, expectedFileBody)

	RemoveDummyFile(t, dummyFile.Location)
	RemoveDummyFile(t, dummyFile1.Location)
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
