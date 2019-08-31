package terra

import (
	"fmt"
	"io/ioutil"
	"os"
)

const (
	AwsDefaultRegion   = "AWS_DEFAULT_REGION"
	AwsRegion          = "AWS_REGION"
	AwsProfile         = "AWS_PROFILE"
	AwsAccessKeyId     = "AWS_ACCESS_KEY_ID"
	AwsSecretAccessKey = "AWS_SECRET_ACCESS_KEY"
)

var (
	DefaultRecipes = map[string]CombinedRecipe{
		"quorum": {
			FilePaths: []string{
				"terra/networking/main.tf",
				"terra/networking/variables.tf",
				"terra/bastion/iam-quorum.tf",
				"terra/bastion/main-quorum.tf",
				"terra/bastion/sg-quorum.tf",
				"terra/quorum/asg.tf",
				"terra/quorum/asg_ecs.tf",
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
				Body: StaticQuorum,
				envVariablesMap: map[string]string{
					"region":                AwsRegion,
					"profile":               AwsProfile,
					"default_region":        AwsDefaultRegion,
					"aws_access_key_id":     AwsAccessKeyId,
					"aws_secret_access_key": AwsSecretAccessKey,
				},
			},
		},
	}
)

type File struct {
	Location             string
	Body                 string
	Variables            map[string]interface{}
	EnvVariablesRollBack map[string]string
	envVariablesMap      map[string]string
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

func (combinedRecipe *CombinedRecipe) BindYamlWithVars(yamlFilePath string) (err error) {
	schema := VariablesSchema{
		File: File{Location: yamlFilePath},
	}

	err = schema.Read()

	if nil != err {
		return err
	}

	if nil == combinedRecipe.Variables {
		combinedRecipe.Variables = make(map[string]interface{}, 0)
	}

	if nil == combinedRecipe.EnvVariablesRollBack {
		combinedRecipe.EnvVariablesRollBack = make(map[string]string, 0)
	}

	for schemaKey, value := range schema.Variables {
		err = combinedRecipe.handleAssignVars(schemaKey, value)

		if nil != err {
			return err
		}
	}

	return err
}

func (combinedRecipe *CombinedRecipe) UnbindEnvVars() (err error) {
	if nil == combinedRecipe.EnvVariablesRollBack {
		return
	}

	for envKey, envVar := range combinedRecipe.EnvVariablesRollBack {
		err = os.Setenv(envKey, envVar)

		if nil != err {
			return err
		}
	}

	return
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

func (combinedRecipe *CombinedRecipe) handleAssignVars(schemaKey string, value interface{}) (err error) {
	combinedRecipe.Variables[schemaKey] = value

	if len(combinedRecipe.envVariablesMap) < 1 {
		return
	}

	envKey := combinedRecipe.envVariablesMap[schemaKey]

	fmt.Printf(
		"\n Trying to assign env variables from recipe %s \n",
		combinedRecipe.envVariablesMap,
	)

	if len(envKey) < 1 {
		fmt.Println("No variables to assign")
		return
	}

	isPreviousRollbackSet := len(combinedRecipe.EnvVariablesRollBack[envKey]) > 0

	if false == isPreviousRollbackSet {
		previousEnv := os.Getenv(envKey)
		combinedRecipe.EnvVariablesRollBack[envKey] = previousEnv
	}

	stringVar := value.(string)
	err = os.Setenv(envKey, stringVar)

	if nil != err {
		return err
	}

	fmt.Printf("\n Assigned env key: %s with value: %s \n", envKey, value)

	return
}
