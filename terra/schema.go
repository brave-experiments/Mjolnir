package terra

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"path"
)

type VariablesSchema struct {
	File
	Type    string
	Version float64
}

type variablesModel struct {
	Version      float64                `yaml:"version"`
	ResourceType string                 `yaml:"resourceType"`
	Variables    map[string]interface{} `yaml:"variables"`
}

const (
	CurrentVersion = float64(0.1)
)

var (
	SupportedFileTypes     = []string{".yml", ".yaml"}
	SupportedResourceTypes = []string{"variables"}
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

	err = yaml.Unmarshal([]byte(variablesSchema.Body), &varsModel)

	if nil != err {
		return err
	}

	err = varsModel.guard()

	if nil != err {
		return err
	}

	variablesSchema.Variables = varsModel.Variables
	variablesSchema.Type = varsModel.ResourceType
	variablesSchema.Version = varsModel.Version

	return err
}

func guardExtension(filePath string) (err error) {
	fileExtension := path.Ext(filePath)

	if false == contains(SupportedFileTypes, fileExtension) {
		return ClientError{fmt.Sprintf(
			"%s is not in supported file types. Valid are: %s",
			fileExtension,
			SupportedFileTypes,
		),
		}
	}

	return
}

func (variablesModel *variablesModel) guard() (err error) {
	err = variablesModel.guardVersion()

	if nil != err {
		return err
	}

	err = variablesModel.guardResourceType()

	if nil != err {
		return err
	}

	err = variablesModel.guardVariables()

	return err
}

func (variablesModel *variablesModel) guardVersion() (err error) {
	version := variablesModel.Version

	if version > CurrentVersion {
		return ClientError{fmt.Sprintf(
			"%v version is not supported. Current version: %v",
			version,
			CurrentVersion,
		),
		}
	}

	return
}

func (variablesModel *variablesModel) guardResourceType() (err error) {
	resourceType := variablesModel.ResourceType

	if false == contains(SupportedResourceTypes, resourceType) {
		return ClientError{fmt.Sprintf(
			"%s is not in supported resource types. Valid are: %s",
			resourceType,
			SupportedResourceTypes,
		),
		}
	}

	return
}

func (variablesModel *variablesModel) guardVariables() (err error) {
	variables := variablesModel.Variables

	if nil == variables {
		return ClientError{"No variables found"}
	}

	return
}

func contains(haystack []string, needle string) bool {
	for _, element := range haystack {
		if needle == element {
			return true
		}
	}

	return false
}
