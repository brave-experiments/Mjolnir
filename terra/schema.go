package terra

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"path"
	"strconv"
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
	err = variablesSchema.ValidateSchemaVariables()

	return err
}

func (variablesSchema *VariablesSchema) ValidateSchemaVariables() (err error) {
	err = variablesSchema.validateClockSkewVariable()

	return err
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

func (variablesSchema *VariablesSchema) validateClockSkewVariable() (err error) {
	if nil == variablesSchema.Variables[ClockSkewKey] {
		return nil
	}

	clockSkewVariables := variablesSchema.Variables[ClockSkewKey].([]interface{})

	for _, variable := range clockSkewVariables {
		err = validateClockSkewVariable(variable)

		if nil != err {
			return err
		}
	}

	return nil
}

func validateClockSkewVariable(variable interface{}) (err error) {
	variableString := variable.(string)
	sign := variableString[:1]
	_, err = strconv.ParseInt(sign, 10, 64)

	if nil == err {
		variableString = "+" + variableString
	}

	isProperSign := contains(SupportedClockSkewSigns, variableString[:1])

	if false == isProperSign {
		return ClientError{
			fmt.Sprintf(
				"%s is not in supported faketime variable signs. Valid are: %s",
				variable,
				SupportedClockSkewSigns,
			),
		}
	}

	variableLength := len(variableString)
	possibleUnit := variableString[variableLength-1:]
	isUnit := contains(SupportedClockSkewUnits, possibleUnit)
	_, err = strconv.ParseInt(possibleUnit, 10, 64)

	if nil != err && false == isUnit {
		return ClientError{
			fmt.Sprintf(
				"%s is not in supported faketime variable units. Valid are: %s",
				variable,
				SupportedClockSkewUnits,
			),
		}
	}

	endIndex := variableLength

	if isUnit {
		endIndex = variableLength - 1
	}

	shouldBeInteger := variableString[1:endIndex]

	_, err = strconv.ParseInt(shouldBeInteger, 10, 64)

	if nil != err {
		return ClientError{"Invalid value, should be integer between sign and faketime unit"}
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
