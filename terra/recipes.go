package terra

import (
	"fmt"
	"io/ioutil"
)

var (
	DefaultRecipes = map[string]CombinedRecipe{
		"quorum": {
			FilePaths: []string{
				"terra/networking/main.tf",
				"terra/bastion/iam-quorum.tf",
				"terra/bastion/main-quorum.tf",
				"terra/bastion/sg-quorum.tf",
				"terra/quorum/asg.tf",
				"terra/quorum/container_definition_bootstrap.tf",
				"terra/quorum/container_definitions.tf",
				"terra/quorum/container_definitions_constellation.tf",
				"terra/quorum/container_definitions_quorum.tf",
				"terra/quorum/container_definitions_tessera.tf",
				"terra/quorum/ecs.tf",
				"terra/quorum/iam.tf",
				"terra/quorum/logging.tf",
				"terra/quorum/main.tf",
				"terra/quorum/outputs.tf",
				"terra/quorum/security_groups.tf",
				"terra/shared/variables.tf",
			},
			File: File{
				Variables: map[string]interface{}{
					"network_name": "sidechain-sandbox",
					"client_name":  "quorum",
					"region":       "us-east-2",
					"profile":      "default",
				},
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
