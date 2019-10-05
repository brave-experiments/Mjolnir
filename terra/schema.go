package terra

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"math/rand"
	"path"
	"strconv"
	"time"
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
	CurrentVersion = float64(0.3)
	NetworkNameKey = "network_name"
	ClockSkewKey   = "faketime"
	SchemaV02      = `version: 0.2
resourceType: variables
variables:
  simpleKey: variable
  region:                     'us-east-2'     ## You can set region for deployment here
  default_region:             'us-east-2'     ## If key region is not present it is default region setter
  profile:                    'default'       ## It chooses profile from your ~/.aws config. If not present, profile is "default"
  network_name:               'sidechain-example'
  number_of_nodes:            '5'
  quorum_docker_image_tag:    '2.2.5'
  aws_access_key_id:          'dummyValue'    ## It overrides access key id env variable. If omitted system env is used
  aws_secret_access_key:      'dummyValue'    ## It overrides secret access key env variable. If omitted system env is used
  genesis_gas_limit:          "28"            ## Used to set genesis gas limit
  is_timestamp:               "30"            ## Used to set genesis timestamp
  genesis_difficulty:         "12"            ## Used to set genesis difficulty
  genesis_nonce:              "0"             ## Used to set genesis nonce
  consensus_mechanism:        "instanbul"     ## Used to set consensus mechanism
  chaos_testing_run_command:  ["netem", "--duration", "5m", "--interface", "eth0", "delay", "--time", "3000", "--jitter", "30", "--correlation", "20", "re2:^ecs-quorum*"]
  faketime:                   ["+2d", "-3h", "+120", "0", "0"]  ## You need to fill all values for existing number of nodes for now.
  tf_log:                     "" ## Used to enable/disable or point logs type that Terraform outputs to console
`
	// Deprecated, remove in 0.3
	SchemaV01 = `version: 0.1
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
	SupportedFileTypes      = []string{".yml", ".yaml"}
	SupportedResourceTypes  = []string{"variables"}
	SupportedClockSkewSigns = []string{"+", "-"}
	SupportedClockSkewUnits = []string{"s", "m", "h", "d", "y"}
)

func (variablesSchema *VariablesSchema) Read() (err error) {
	err = variablesSchema.guard()

	if nil != err {
		return err
	}

	variablesSchema.mapGenesisVariables()
	variablesSchema.mapNetworkName()

	return
}

func (variablesSchema *VariablesSchema) mapNetworkName() {
	variables := variablesSchema.Variables

	if nil == variables {
		return
	}

	networkName := variables[NetworkNameKey]

	if nil == networkName {
		return
	}

	fixedSalt := rand.New(rand.NewSource(time.Now().UnixNano()))
	randomString := fmt.Sprintf("%v", fixedSalt.Int())
	variablesSchema.Variables[NetworkNameKey] = fmt.Sprintf("%s-%v", networkName, randomString[0:8])
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

	err = variablesSchema.validateClockSkewVariable()

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

func (variablesSchema *VariablesSchema) validateClockSkewVariable() (err error) {
	if nil != variablesSchema.Variables[ClockSkewKey] {
		clockSkewVariables := variablesSchema.Variables[ClockSkewKey].([]interface{})

		for index, variable := range clockSkewVariables {
			variableString := variable.(string)
			variableLength := len(variableString)
			sign := variableString[:1]
			unit := variableString[variableLength-1:]

			if false == contains(SupportedClockSkewUnits, unit) {
				if "" != unit && "0" != unit {
					return ClientError{fmt.Sprintf(
							"%s is not in supported faketime variable units. Valid are: %s",
							variable,
							SupportedClockSkewUnits,
						),
					}
				}
			}

			if false == contains(SupportedClockSkewSigns, sign) {
				if "" != variableString[:variableLength-1] {
					if _, err := strconv.ParseInt(variableString,10,64); nil != err {
						variableString = variableString[:variableLength-1]
					}
				}

				intVariable, err := strconv.ParseInt(variableString, 10, 0)

				if nil != err {
					return err
				}

				if 0 == intVariable {
					continue
				}

				if 0 < intVariable {
					variable = SupportedClockSkewSigns[0] + variableString
					variablesSchema.Variables[ClockSkewKey].([]interface{})[index] = variable
					continue
				}

				return ClientError{fmt.Sprintf(
						"%s is not in supported faketime variable signs. Valid are: %s",
						variable,
						SupportedClockSkewSigns,
					),
				}
			}

			_, err := strconv.ParseInt(variableString[1:variableLength-1], 10, 0)

			if nil != err {
				return err
			}
		}
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
