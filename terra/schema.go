package terra

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"path"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

var (
	VariablesKeyToHex = []string{
		GenesisGasLimitKey,
		"genesis_timestamp",
		"genesis_difficulty",
		"genesis_nonce",
		"genesis_blocktime",
		GenesisMinGasLimitKey,
	}
	ValidRegions = []string{
		"us-east-1",
		"us-east-2",
		"us-west-1",
		"us-west-2",
		"ca-central-1",
		"eu-central-1",
		"eu-west-1",
		"eu-west-2",
		"eu-west-3",
		"eu-north-1",
		"ap-east-1",
		"ap-northeast-1",
		"ap-northeast-2",
		"ap-northeast-3",
		"ap-southeast-1",
		"ap-southeast-2",
		"ap-south-1",
		"me-south-1",
		"sa-east-1",
	}
	kValidInstances = map[string]awsInstance{
		"a1.medium":     awsInstance{Type: "a1.medium", Memory: "8192", Cpu: "4096"},
		"a1.large":      awsInstance{Type: "a1.large", Memory: "8192", Cpu: "4096"},
		"a1.xlarge":     awsInstance{Type: "a1.xlarge", Memory: "8192", Cpu: "4096"},
		"a1.2xlarge":    awsInstance{Type: "a1.2xlarge", Memory: "8192", Cpu: "4096"},
		"a1.4xlarge":    awsInstance{Type: "a1.4xlarge", Memory: "8192", Cpu: "4096"},
		"t3.nano":       awsInstance{Type: "t3.nano", Memory: "8192", Cpu: "4096"},
		"t3.micro":      awsInstance{Type: "t3.micro", Memory: "8192", Cpu: "4096"},
		"t3.small":      awsInstance{Type: "t3.small", Memory: "8192", Cpu: "4096"},
		"t3.medium":     awsInstance{Type: "t3.medium", Memory: "8192", Cpu: "4096"},
		"t3.large":      awsInstance{Type: "t3.large", Memory: "8192", Cpu: "4096"},
		"t3.xlarge":     awsInstance{Type: "t3.xlarge", Memory: "8192", Cpu: "4096"},
		"t3.2xlarge":    awsInstance{Type: "t3.2xlarge", Memory: "8192", Cpu: "4096"},
		"t3a.nano":      awsInstance{Type: "t3a.nano", Memory: "8192", Cpu: "4096"},
		"t3a.micro":     awsInstance{Type: "t3a.micro", Memory: "8192", Cpu: "4096"},
		"t3a.small":     awsInstance{Type: "t3a.small", Memory: "8192", Cpu: "4096"},
		"t3a.medium":    awsInstance{Type: "t3a.medium", Memory: "8192", Cpu: "4096"},
		"t3a.large":     awsInstance{Type: "t3a.large", Memory: "8192", Cpu: "4096"},
		"t3a.xlarge":    awsInstance{Type: "t3a.xlarge", Memory: "8192", Cpu: "4096"},
		"t3a.2xlarge":   awsInstance{Type: "t3a.2xlarge", Memory: "8192", Cpu: "4096"},
		"t2.nano":       awsInstance{Type: "t2.nano", Memory: "8192", Cpu: "4096"},
		"t2.micro":      awsInstance{Type: "t2.micro", Memory: "8192", Cpu: "4096"},
		"t2.small":      awsInstance{Type: "t2.small", Memory: "8192", Cpu: "4096"},
		"t2.medium":     awsInstance{Type: "t2.medium", Memory: "8192", Cpu: "4096"},
		"t2.large":      awsInstance{Type: "t2.large", Memory: "8192", Cpu: "4096"},
		"t2.xlarge":     awsInstance{Type: "t2.xlarge", Memory: "16384", Cpu: "4096"},
		"t2.2xlarge":    awsInstance{Type: "t2.2xlarge", Memory: "8192", Cpu: "4096"},
		"m5.large":      awsInstance{Type: "m5.large", Memory: "8192", Cpu: "4096"},
		"m5.xlarge":     awsInstance{Type: "m5.xlarge", Memory: "8192", Cpu: "4096"},
		"m5.2xlarge":    awsInstance{Type: "m5.2xlarge", Memory: "8192", Cpu: "4096"},
		"m5.4xlarge":    awsInstance{Type: "m5.4xlarge", Memory: "8192", Cpu: "4096"},
		"m5.8xlarge":    awsInstance{Type: "m5.8xlarge", Memory: "8192", Cpu: "4096"},
		"m5.12xlarge":   awsInstance{Type: "m5.12xlarge", Memory: "8192", Cpu: "4096"},
		"m5.16xlarge":   awsInstance{Type: "m5.16xlarge", Memory: "8192", Cpu: "4096"},
		"m5.24xlarge":   awsInstance{Type: "m5.24xlarge", Memory: "8192", Cpu: "4096"},
		"m5.metal":      awsInstance{Type: "m5.metal", Memory: "8192", Cpu: "4096"},
		"m5d.large":     awsInstance{Type: "m5d.large", Memory: "8192", Cpu: "4096"},
		"m5d.xlarge":    awsInstance{Type: "m5d.xlarge", Memory: "8192", Cpu: "4096"},
		"m5d.2xlarge":   awsInstance{Type: "m5d.2xlarge", Memory: "8192", Cpu: "4096"},
		"m5d.4xlarge":   awsInstance{Type: "m5d.4xlarge", Memory: "8192", Cpu: "4096"},
		"m5d.8xlarge":   awsInstance{Type: "m5d.8xlarge", Memory: "8192", Cpu: "4096"},
		"m5d.12xlarge":  awsInstance{Type: "m5d.12xlarge", Memory: "8192", Cpu: "4096"},
		"m5d.16xlarge":  awsInstance{Type: "m5d.16xlarge", Memory: "8192", Cpu: "4096"},
		"m5d.24xlarge":  awsInstance{Type: "m5d.24xlarge", Memory: "8192", Cpu: "4096"},
		"m5d.metal":     awsInstance{Type: "m5d.metal", Memory: "8192", Cpu: "4096"},
		"m5a.large":     awsInstance{Type: "m5a.large", Memory: "8192", Cpu: "4096"},
		"m5a.xlarge":    awsInstance{Type: "m5a.xlarge", Memory: "8192", Cpu: "4096"},
		"m5a.2xlarge":   awsInstance{Type: "m5a.2xlarge", Memory: "8192", Cpu: "4096"},
		"m5a.4xlarge":   awsInstance{Type: "m5a.4xlarge", Memory: "8192", Cpu: "4096"},
		"m5a.8xlarge":   awsInstance{Type: "m5a.8xlarge", Memory: "8192", Cpu: "4096"},
		"m5a.12xlarge":  awsInstance{Type: "m5a.12xlarge", Memory: "8192", Cpu: "4096"},
		"m5a.16xlarge":  awsInstance{Type: "m5a.16xlarge", Memory: "8192", Cpu: "4096"},
		"m5a.24xlarge":  awsInstance{Type: "m5a.24xlarge", Memory: "8192", Cpu: "4096"},
		"m5ad.large":    awsInstance{Type: "m5ad.large", Memory: "8192", Cpu: "4096"},
		"m5ad.xlarge":   awsInstance{Type: "m5ad.xlarge", Memory: "8192", Cpu: "4096"},
		"m5ad.2xlarge":  awsInstance{Type: "m5ad.2xlarge", Memory: "8192", Cpu: "4096"},
		"m5ad.4xlarge":  awsInstance{Type: "m5ad.4xlarge", Memory: "8192", Cpu: "4096"},
		"m5ad.12xlarge": awsInstance{Type: "m5ad.12xlarge", Memory: "8192", Cpu: "4096"},
		"m5ad.24xlarge": awsInstance{Type: "m5ad.24xlarge", Memory: "8192", Cpu: "4096"},
		"m4.large":      awsInstance{Type: "m4.large", Memory: "8192", Cpu: "4096"},
		"m4.xlarge":     awsInstance{Type: "m4.xlarge", Memory: "8192", Cpu: "4096"},
		"m4.2xlarge":    awsInstance{Type: "m4.2xlarge", Memory: "8192", Cpu: "4096"},
		"m4.4xlarge":    awsInstance{Type: "m4.4xlarge", Memory: "8192", Cpu: "4096"},
		"m4.10xlarge":   awsInstance{Type: "m4.10xlarge", Memory: "8192", Cpu: "4096"},
		"m4.16xlarge":   awsInstance{Type: "m4.16xlarge", Memory: "8192", Cpu: "4096"},
	}
	ValidInstanceTypes        []string
	ValidConsensusMechanisms  = []string{"raft", "instanbul"}
	SupportedFileTypes        = []string{".yml", ".yaml"}
	SupportedResourceTypes    = []string{"variables"}
	SupportedClockSkewSigns   = []string{"+", "-"}
	SupportedClockSkewUnits   = []string{"s", "m", "h", "d", "y"}
	StringVariablesToValidate = []string{"region", "default_region", "profile", "aws_access_key_id", "aws_secret_access_key"}
)

func init() {
	// Initialize the list of valid instance types by taking the keys
	// from the list of valid instances
	ValidInstanceTypes = make([]string, len(ValidInstances))
	i := 0
	for key := range ValidInstances {
		ValidInstanceTypes[i] = key
		i++
	}
}

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

type awsInstance struct {
	Type   string
	Memory string
	Cpu    string
}

const (
	CurrentVersion          = float64(0.3)
	MaxNetworkNameVarLength = 20
	GenesisMinGasLimitKey   = "genesis_min_gas_limit"
	GenesisGasLimitKey      = "genesis_gas_limit"
	NetworkNameKey          = "network_name"
	ClockSkewKey            = "faketime"
	NodeNumbersKey          = "number_of_nodes"
	QuorumDockerImageTagKey = "quorum_docker_image_tag"
	NetworkNameTerraRegExp  = "[a-z]([-a-z0-9]*[a-z0-9])?"
	SchemaV02               = `version: 0.2
resourceType: variables
variables:
  simpleKey: variable
  region:                     'us-east-2'     ## You can set region for deployment here
  default_region:             'us-east-2'     ## If key region is not present it is default region setter
  profile:                    'default'       ## It chooses profile from your ~/.aws config. If not present, profile is "default"
  network_name:               'sidechain'
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

func (variablesSchema *VariablesSchema) Read() (err error) {
	err = variablesSchema.guard()
	if nil != err {
		return yamlValidationError(err.Error())
	}

	err = variablesSchema.mapGenesisVariables()
	if nil != err {
		return yamlValidationError(err.Error())
	}

	// Dynamically calculate and set ecs_cpu and ecs_memory
	// depending on the asg_instance_type supplied in values yaml
	err = variablesSchema.mapAwsInstanceMemoryAndCpu()
	if nil != err {
		return yamlValidationError(err.Error())
	}

	err = variablesSchema.ValidateSchemaVariables()

	return err
}

func (variablesSchema *VariablesSchema) ValidateSchemaVariables() (err error) {
	err = variablesSchema.validateClockSkewVariable()

	if nil != err {
		return err
	}

	err = variablesSchema.validateNodeNumbersVariable()

	if nil != err {
		return err
	}

	err = variablesSchema.validateQuorumDockerImageTagVariable()

	if nil != err {
		return err
	}

	err = variablesSchema.validateNetworkNameVariable()

	if nil != err {
		return err
	}

	err = variablesSchema.validateNonSpacesStringVariable()

	if nil != err {
		return err
	}

	err = variablesSchema.validateAwsProperties()

	if nil != err {
		return err
	}

	err = variablesSchema.validateConsensus()

	if nil != err {
		return err
	}

	return nil
}

func (variablesSchema *VariablesSchema) mapGenesisVariables() (err error) {
	for key, variable := range variablesSchema.Variables {
		if false == contains(VariablesKeyToHex, key) {
			continue
		}

		hexValue, err := ConvertInterfaceToHex(variable)

		if nil != err {
			return err
		}

		variablesSchema.Variables[key] = hexValue
	}

	err = variablesSchema.guardMinBlockGasLimit()

	return err
}

func (variablesSchema *VariablesSchema) mapAwsInstanceMemoryAndCpu() (err error) {
	// Look up instance information in the Variables
	var desiredInstanceType string
	for key, variable := range variablesSchema.Variables {
		if key == "asg_instance_type" {
			desiredInstanceType = variable.(string)
		}
	}

	desiredInstance := ValidInstances[desiredInstanceType]
	if desiredInstance.Memory == "" || desiredInstance.Cpu == "" {
		// FIXME
		// Return a ClientError instead of silently using defaults
		return nil
	}

	variablesSchema.Variables["ecs_memory"] = desiredInstance.Memory
	variablesSchema.Variables["ecs_cpu"] = desiredInstance.Cpu

	return nil
}

func (variablesSchema *VariablesSchema) guardMinBlockGasLimit() (err error) {
	variables := variablesSchema.Variables

	if nil == variables {
		return nil
	}

	genesisGasLimit, gasLimitFound := variables[GenesisGasLimitKey]
	genesisMinGasLimit, minGasLimitFound := variables[GenesisMinGasLimitKey]

	if false == gasLimitFound || false == minGasLimitFound {
		return nil
	}

	if isHexGreaterThanOrEqual(genesisGasLimit, genesisMinGasLimit) {
		errorMessage := fmt.Sprintf(
			"%s must be greater than %s",
			GenesisMinGasLimitKey,
			GenesisGasLimitKey,
		)

		return ClientError{Message: errorMessage}
	}

	return nil
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
		return ClientError{
			fmt.Sprintf(
				"%s is not in supported file types. Valid are: %s",
				fileExtension,
				SupportedFileTypes,
			),
		}
	}

	return nil
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

	return nil
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

	return nil
}

func (variablesModel *variablesModel) guardVariables() (err error) {
	variables := variablesModel.Variables

	if nil == variables {
		return ClientError{"No variables found"}
	}

	return nil
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

func (variablesSchema *VariablesSchema) validateConsensus() (err error) {
	desiredKey := "consensus_mechanism"

	value, ok := variablesSchema.Variables[desiredKey]

	if false == ok {
		return nil
	}

	if contains(ValidConsensusMechanisms, value.(string)) {
		return nil
	}

	return ClientError{
		fmt.Sprintf("Invalid %s, valid are: %s", desiredKey, ValidConsensusMechanisms),
	}
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

func (variablesSchema *VariablesSchema) validateNodeNumbersVariable() (err error) {
	if nil == variablesSchema.Variables[NodeNumbersKey] {
		return nil
	}

	nodeNumbersVariable := variablesSchema.Variables[NodeNumbersKey].(string)
	_, err = strconv.ParseInt(nodeNumbersVariable, 10, 64)

	if nil != err {
		return ClientError{
			fmt.Sprintf(
				"%s is not in supported node of numbers variable value type.",
				nodeNumbersVariable,
			),
		}
	}

	return nil
}

func (variablesSchema *VariablesSchema) validateQuorumDockerImageTagVariable() (err error) {
	if nil == variablesSchema.Variables[QuorumDockerImageTagKey] {
		return nil
	}

	nodeNumbersVariable := variablesSchema.Variables[QuorumDockerImageTagKey].(string)

	versionIntegers := strings.Split(nodeNumbersVariable, ".")

	for _, intVal := range versionIntegers {
		_, err = strconv.ParseInt(intVal, 10, 64)

		if nil != err {
			return ClientError{"Invalid value, should be integer docker image tag version between dots"}
		}
	}

	return nil
}

func (variablesSchema *VariablesSchema) validateNetworkNameVariable() (err error) {
	if nil == variablesSchema.Variables[NetworkNameKey] {
		return nil
	}

	networkNameVariable := variablesSchema.Variables[NetworkNameKey].(string)

	if MaxNetworkNameVarLength < len(networkNameVariable) {
		return ClientError{
			fmt.Sprintf(
				"Network name is too long. Maximum allowed characters are %d",
				MaxNetworkNameVarLength,
			),
		}
	}

	variableDeclarationRegex := regexp.MustCompile(NetworkNameTerraRegExp)
	variableDeclarations := variableDeclarationRegex.FindString(networkNameVariable)

	if networkNameVariable != variableDeclarations {
		return ClientError{"Network name is invalid"}
	}

	return nil
}

func (variablesSchema *VariablesSchema) validateNonSpacesStringVariable() (err error) {
	if nil == variablesSchema.Variables {
		return nil
	}

	for key, variable := range variablesSchema.Variables {
		if false != contains(StringVariablesToValidate, key) {
			err = validateNonSpacesStringVariable(variable)

			if nil != err {
				return ClientError{
					fmt.Sprintf(
						"Variable with key: %s contains white space which is not allowed",
						key,
					),
				}
			}
		}
	}

	return nil
}

func (variablesSchema *VariablesSchema) validateAwsProperties() (err error) {
	if nil == variablesSchema.Variables {
		return nil
	}

	for key, variable := range variablesSchema.Variables {
		err = validateAwsRegions(key, variable)
		if nil != err {
			return err
		}

		err = validateAwsInstanceType(key, variable)
		if nil != err {
			return err
		}
	}

	return nil
}

func validateAwsRegions(key string, variable interface{}) (err error) {
	desiredKeys := []string{"region", "default_region"}

	if false == contains(desiredKeys, key) {
		return nil
	}

	stringVariable := fmt.Sprintf("%v", variable)

	if contains(ValidRegions, stringVariable) {
		return nil
	}

	return ClientError{
		fmt.Sprintf(
			"%s is not valid AWS region. Valid are: %s",
			stringVariable,
			ValidRegions,
		),
	}
}

func validateAwsInstanceType(key string, variable interface{}) (err error) {
	desiredKey := "asg_instance_type"

	if key != desiredKey {
		return nil
	}

	stringVariable := fmt.Sprintf("%v", variable)
	if contains(ValidInstanceTypes, stringVariable) {
		return nil
	}

	return ClientError{
		fmt.Sprintf(
			"%s is not valid %s, valid are: %s",
			stringVariable,
			desiredKey,
			ValidInstanceTypes,
		),
	}
}

func validateNonSpacesStringVariable(variable interface{}) (err error) {
	for _, v := range variable.(string) {
		if unicode.IsSpace(v) {
			return ClientError{}
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

func yamlValidationError(body string) (clientError ClientError) {
	return ClientError{
		Message: fmt.Sprintf("\n[ERR] Yaml Validation error: %s", body),
	}
}
