package terra

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"path"
)

var (
	VariablesKeyToHex = []string{
		"genesis_gas_limit",
		"genesis_timestamp",
		"genesis_difficulty",
		"genesis_nonce",
	}
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
	SchemaV1       = `version: 0.1
resourceType: variables
variables:
  simpleKey: variable
  region:                'us-east-1'     ## You can set region for deployment here
  default_region:        'us-west-1'     ## If key region is not present it is default region setter
  profile:               'default'       ## It chooses profile from your ~/.aws config. If not present, profile is "default"
  aws_access_key_id:     'dummyValue'    ## It overrides access key id env variable. If omitted system env is used
  aws_secret_access_key: 'dummyValue'    ## It overrides secret access key env variable. If omitted system env is used
  genesis_gas_limit:      25		     ## Used to set genesis gas limit it converts to hex
  genesis_timestamp:      38	         ## Used to set genesis timestamp it converts to hex
  genesis_difficulty:     12             ## Used to set genesis difficulty it converts to hex
  genesis_nonce:          0              ## Used to set genesis nonce it converts to hex
  consensus_mechanism:    "instanbul"    ## Used to set consensus mechanism supported values are raft/istanbul
`
)

var (
	SupportedFileTypes     = []string{".yml", ".yaml"}
	SupportedResourceTypes = []string{"variables"}
)

func (variablesSchema *VariablesSchema) Read() (err error) {
	err = variablesSchema.guard()

	if nil != err {
		return err
	}

	variablesSchema.mapGenesisVariables()

	return
}

func (variablesSchema *VariablesSchema) mapGenesisVariables() {
	for key, variable := range variablesSchema.Variables {
		if false == contains(VariablesKeyToHex, key) {
			continue
		}

		variablesSchema.Variables[key] = ConvertInterfaceToHex(variable)
	}
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
