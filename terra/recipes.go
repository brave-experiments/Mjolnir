package terra

import (
	"fmt"
	"io/ioutil"
)

var (
	DefaultRecipes = map[string]File{
		"bastion": {
			Location: "bastion.tf",
		},
	}
)

type File struct {
	Location string
	Body     string
}

type Recipes struct {
	Elements map[string]File
}

type RecipesError struct {
	Message string
}

func (recipesError RecipesError) Error() string {
	return recipesError.Message
}

func (recipes *Recipes) CreateWithDefaults() {
	recipes.Elements = DefaultRecipes
}

func (recipes *Recipes) AddRecipe(keyName string, file File) error {
	if nil == recipes.Elements {
		recipes.Elements = make(map[string]File, 0)
	}

	if _, ok := recipes.Elements[keyName]; ok {
		return RecipesError{fmt.Sprintf("%s  already exists in recipes list", keyName)}
	}

	recipes.Elements[keyName] = file

	return nil
}

func (file *File) ReadFile() error {
	fileBodyBytes, err := ioutil.ReadFile(file.Location)

	if nil != err {
		return err
	}

	file.Body = string(fileBodyBytes)

	return nil
}
