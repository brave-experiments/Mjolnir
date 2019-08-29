package terra

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"path"
)

type VariablesSchema struct {
	File
	Type string
}

type variablesModel struct {
	ApolloSchema string                 `yaml:"apolloSchema"`
	ResourceType string                 `yaml:"resourceType"`
	Variables    map[string]interface{} `yaml:"variables"`
}

var (
	SupportedTypes = []string{".yml", ".yaml"}
)

//func (variablesSchema *VariablesSchema) ReadFile() (err error) {
//
//}

func (variablesSchema *VariablesSchema) Read() (err error) {
	err = variablesSchema.guard()

	return err
}

func (variablesSchema *VariablesSchema) guard() (err error) {
	err = variablesSchema.ReadFile()

	if nil != err {
		return err
	}

	err = guardExtension(variablesSchema.Location)

	if nil != err {
		return err
	}

	varsModel := variablesModel{}

	err = yaml.Unmarshal([]byte(variablesSchema.Body), varsModel)

	return err
}

func guardExtension(filePath string) (err error) {
	fileExtension := path.Ext(filePath)

	if false == contains(SupportedTypes, fileExtension) {
		return ClientError{fmt.Sprintf(
			"%s is not in supported types. Valid are: %s",
			fileExtension,
			SupportedTypes,
		),
		}
	}

	return nil
}

func contains(haystack []string, needle string) bool {
	for _, element := range haystack {
		if needle == element {
			return true
		}
	}
	return false
}
