package terra

import (
	"fmt"
	"io/ioutil"
)

var (
	DefaultRecipes = map[string]Recipe{
		"bastion": {
			Location: "bastion.tf",
		},
	}
)

type Recipe struct {
	Location string
	Body     string
}

type Recipes struct {
	Elements map[string]Recipe
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

func (recipes *Recipes) AddRecipe(keyName string, recipe Recipe) error {
	if nil == recipes.Elements {
		recipes.Elements = make(map[string]Recipe, 0)
	}

	if _, ok := recipes.Elements[keyName]; ok {
		return RecipesError{fmt.Sprintf("%s  already exists in recipes list", keyName)}
	}

	recipes.Elements[keyName] = recipe

	return nil
}

func (recipe *Recipe) Create() error {
	fileBodyBytes, err := ioutil.ReadFile(recipe.Location)

	if nil != err {
		return err
	}

	recipe.Body = string(fileBodyBytes)

	return nil
}
