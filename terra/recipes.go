package terra

import (
	"fmt"
	"io/ioutil"
)

var (
	DefaultRecipes = map[string]CombinedRecipe{
		"bastion": {
			FilePaths: []string{
				"bastion.tf",
			},
		},
	}
)

type File struct {
	Location  string
	Body      string
	Variables map[string]interface{}
}

type CombinedRecipe struct {
	File
	FilePaths []string
}

type Recipes struct {
	Elements map[string]CombinedRecipe
}

type RecipesError struct {
	Message string
}

type ClientError struct {
	Message string
}

func (recipesError RecipesError) Error() string {
	return recipesError.Message
}

func (client ClientError) Error() string {
	return client.Message
}

func (recipes *Recipes) CreateWithDefaults() {
	recipes.Elements = DefaultRecipes
}

func (recipes *Recipes) AddRecipe(keyName string, combinedRecipe CombinedRecipe) error {
	if nil == recipes.Elements {
		recipes.Elements = make(map[string]CombinedRecipe, 0)
	}

	if _, ok := recipes.Elements[keyName]; ok {
		return RecipesError{fmt.Sprintf("%s already exists in recipes list", keyName)}
	}

	recipes.Elements[keyName] = combinedRecipe

	return nil
}

func (combinedRecipe *CombinedRecipe) ParseBody() (err error) {
	filePaths := combinedRecipe.FilePaths

	if nil == filePaths || len(filePaths) < 1 {
		return RecipesError{"There are no recipes within this combined recipe"}
	}

	combinedRecipe.Body = ""

	for _, filePath := range filePaths {
		file := File{
			Location:  filePath,
			Variables: combinedRecipe.Variables,
		}
		err = file.ReadFile()

		if nil != err {
			return err
		}

		combinedRecipe.Body = combinedRecipe.Body + "\n" + file.Body
	}

	return nil
}

func (file *File) ReadFile() (err error) {
	fileBodyBytes, err := ioutil.ReadFile(file.Location)

	if nil != err {
		return err
	}

	file.Body = string(fileBodyBytes)

	return nil
}

func (file *File) WriteFile() (err error) {
	return ioutil.WriteFile(file.Location, []byte(file.Body), 0644)
}
