package terra

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"path"
	"testing"
)

func TestValidInstanceTypes(t *testing.T) {
	assert.Greater(t, len(ValidInstanceTypes), 0)
}

func TestVariablesSchema_ReadFailure(t *testing.T) {
	variablesSchema := VariablesSchema{}
	err := variablesSchema.Read()
	assert.Error(t, err)
	assert.Equal(
		t,
		fmt.Sprintf(NoSuchFileOrDirectoryMsg, variablesSchema.Location),
		err.Error(),
	)

	variablesSchema.Location = "non-existing.yml"
	err = variablesSchema.Read()
	assert.Error(t, err)
	assert.Equal(
		t,
		fmt.Sprintf(NoSuchFileOrDirectoryMsg, variablesSchema.Location),
		err.Error(),
	)

	// It fails on invalid file type
	dummyFilePath := "dummy.tf"
	PrepareDummyFile(t, dummyFilePath, "")
	variablesSchema.Location = dummyFilePath
	err = variablesSchema.Read()
	assert.Error(t, err)
	assert.IsType(t, ClientError{}, err)
	assert.Equal(
		t,
		fmt.Sprintf(NotValidExtMsg, path.Ext(dummyFilePath)),
		err.Error(),
	)
	RemoveDummyFile(t, dummyFilePath)

	// It fails on invalid file body
	dummyFilePath = "dummy.yml"
	PrepareDummyFile(t, dummyFilePath, "{Some string:\n\t\tkk:\nx}")
	variablesSchema.Location = dummyFilePath
	err = variablesSchema.Read()
	assert.Error(t, err)
	assert.Equal(t, "\n[ERR] Yaml Validation error: yaml: line 1: did not find expected ',' or '}'", err.Error())
	RemoveDummyFile(t, dummyFilePath)

	// It fails on invalid file path
	dummyFilePath = "dummyInvalidPath.yml"
	variablesSchema.Location = dummyFilePath
	err = variablesSchema.Read()
	assert.Error(t, err)
	assert.Equal(
		t,
		fmt.Sprintf("\n[ERR] Yaml Validation error: open %s: no such file or directory", dummyFilePath),
		err.Error(),
	)
}

func TestVariablesSchema_ReadFailure_BodyParsing(t *testing.T) {
	// It fails on invalid version
	variablesSchema := VariablesSchema{}
	dummyFilePath := "dummy.yml"
	PrepareDummyFile(t, dummyFilePath, "{Some string:\n\t\tkk:\nx}")
	variablesSchema.Location = dummyFilePath
	err := variablesSchema.Read()
	assert.Error(t, err)
	assert.Equal(t, "\n[ERR] Yaml Validation error: yaml: line 1: did not find expected ',' or '}'", err.Error())
	RemoveDummyFile(t, dummyFilePath)

	// It fails on invalid resource version
	version := float64(2)
	resource := "dummyResource"
	configuredYaml := configureYaml(version, resource)
	PrepareDummyFile(t, dummyFilePath, configuredYaml)
	err = variablesSchema.Read()
	assert.Error(t, err)
	assert.Equal(
		t,
		fmt.Sprintf("\n[ERR] Yaml Validation error: %v version is not supported. Current version: %v", version, CurrentVersion),
		err.Error(),
	)

	// It fails on invalid resource type
	version = float64(0.1)
	configuredYaml = configureYaml(version, resource)
	PrepareDummyFile(t, dummyFilePath, configuredYaml)
	err = variablesSchema.Read()
	assert.Error(t, err)
	assert.Equal(
		t,
		fmt.Sprintf("\n[ERR] Yaml Validation error: %s is not in supported resource types. Valid are: %s", resource, SupportedResourceTypes),
		err.Error(),
	)
	RemoveDummyFile(t, dummyFilePath)

	// It fails when no variables are present
	version = float64(0.1)
	PrepareDummyFile(t, dummyFilePath, YamlFixtureNoVariables)
	err = variablesSchema.Read()
	assert.Error(t, err)
	assert.Equal(
		t,
		"\n[ERR] Yaml Validation error: No variables found",
		err.Error(),
	)
	RemoveDummyFile(t, dummyFilePath)
}

func TestVariablesSchema_Read(t *testing.T) {
	variablesSchema := VariablesSchema{}
	dummyFilePath := "dummy.yml"
	PrepareDummyFile(t, dummyFilePath, YamlV01Fixture)
	variablesSchema.Location = dummyFilePath
	err := variablesSchema.Read()
	assert.Nil(t, err)
	assert.Equal(t, float64(0.1), variablesSchema.Version)
	assert.Equal(t, variablesSchema.Type, "variables")
	assert.NotNil(t, variablesSchema.Variables)
	variables := variablesSchema.Variables
	assert.Equal(t, "variable", variables["simpleKey"])
	RemoveDummyFile(t, dummyFilePath)
}

func TestVariablesSchema_Read_v02(t *testing.T) {
	// Network name should have added hash
	variablesSchema := VariablesSchema{}
	dummyFilePath := "dummy.yml"
	PrepareDummyFile(t, dummyFilePath, YamlV02Fixture)
	variablesSchema.Location = dummyFilePath
	err := variablesSchema.Read()
	assert.Nil(t, err)
	assert.Equal(t, float64(0.2), variablesSchema.Version)
	assert.Equal(t, variablesSchema.Type, "variables")
	assert.NotNil(t, variablesSchema.Variables)
	variables := variablesSchema.Variables
	// Network name should look like `variable-[8-length-random-integer]`
	newNetworkName := variables["network_name"].(string)
	assert.Equal(t, 8, len(newNetworkName))
	RemoveDummyFile(t, dummyFilePath)
}

func TestVariablesSchema_Read_WithHexUtil(t *testing.T) {
	variablesSchema := VariablesSchema{}
	dummyFilePath := "dummy.yml"
	PrepareDummyFile(t, dummyFilePath, YamlFixtureWithHexUtils)
	variablesSchema.Location = dummyFilePath
	err := variablesSchema.Read()
	assert.Nil(t, err)

	assert.Equal(t, 6, len(VariablesKeyToHex))
	assert.Equal(t, "0x19", variablesSchema.Variables[VariablesKeyToHex[0]])
	assert.Equal(t, "0x26", variablesSchema.Variables[VariablesKeyToHex[1]])
	assert.Equal(t, "0xc", variablesSchema.Variables[VariablesKeyToHex[2]])
	assert.Equal(t, "0x0", variablesSchema.Variables[VariablesKeyToHex[3]])

	RemoveDummyFile(t, dummyFilePath)

	// Should throw no error when only one of the genesis comparision are present
	PrepareDummyFile(t, dummyFilePath, YamlFixtureGasLimitWithoutMinGas)
	err = variablesSchema.Read()
	assert.Nil(t, err)
	RemoveDummyFile(t, dummyFilePath)

	// should not fail if one is higher than another
	PrepareDummyFile(t, dummyFilePath, YamlFixtureGasLimitLowetHanMinGasLimit)
	err = variablesSchema.Read()
	assert.Nil(t, err)
	RemoveDummyFile(t, dummyFilePath)
}

func TestVariablesSchema_Read_WithHexUtil_Failure(t *testing.T) {
	variablesSchema := VariablesSchema{}
	dummyFilePath := "dummy.yml"
	PrepareDummyFile(t, dummyFilePath, YamlFixtureWithInvalidHexUtils)
	variablesSchema.Location = dummyFilePath
	err := variablesSchema.Read()
	assert.Error(t, err)
	RemoveDummyFile(t, dummyFilePath)

	// should fail 'cos gas limit cannot be higher than min gas limit
	PrepareDummyFile(t, dummyFilePath, YamlFixtureGasLimitGreaterThanMinGasLimit)
	err = variablesSchema.Read()
	assert.Error(t, err)
	assert.Equal(t, "\n[ERR] Yaml Validation error: genesis_min_gas_limit must be greater than genesis_gas_limit", err.Error())
	RemoveDummyFile(t, dummyFilePath)
}

func TestVariablesSchema_ValidateSchemaVariables(t *testing.T) {
	variablesSchema := VariablesSchema{}
	dummyFilePath := "dummy.yml"
	PrepareDummyFile(t, dummyFilePath, YamlV03Fixture)
	variablesSchema.Location = dummyFilePath
	err := variablesSchema.Read()
	assert.Nil(t, err)
	err = variablesSchema.ValidateSchemaVariables()
	assert.Nil(t, err)

	RemoveDummyFile(t, dummyFilePath)
}

func TestVariablesSchema_Readv03Failure(t *testing.T) {
	variablesSchema := VariablesSchema{}
	dummyFilePath := "dummy.yml"

	PrepareDummyFile(t, dummyFilePath, IncorrectSignYamlV03Fixture)
	variablesSchema.Location = dummyFilePath
	err := variablesSchema.Read()
	assert.Error(t, err)
	assert.Equal(t, ClientError{Message: "@2s is not in supported faketime variable signs. Valid are: [+ -]"}, err)

	PrepareDummyFile(t, dummyFilePath, IncorrectUnitYamlV03Fixture)
	variablesSchema.Location = dummyFilePath
	err = variablesSchema.Read()
	assert.Error(t, err)
	assert.Equal(t, ClientError{Message: "+2x is not in supported faketime variable units. Valid are: [s m h d y]"}, err)

	PrepareDummyFile(t, dummyFilePath, IncorrectValueYamlV03Fixture)
	variablesSchema.Location = dummyFilePath
	err = variablesSchema.Read()
	assert.Error(t, err)
	assert.Equal(t, ClientError{Message: "Invalid value, should be integer between sign and faketime unit"}, err)

	PrepareDummyFile(t, dummyFilePath, IncorrectNodesNumberYamlV03Fixture)
	variablesSchema.Location = dummyFilePath
	err = variablesSchema.Read()
	assert.Error(t, err)
	assert.Equal(t, ClientError{Message: "Abc is not in supported node of numbers variable value type."}, err)

	PrepareDummyFile(t, dummyFilePath, IncorrectQuorumDockerImageTagYamlV03Fixture)
	variablesSchema.Location = dummyFilePath
	err = variablesSchema.Read()
	assert.Error(t, err)
	assert.Equal(t, ClientError{Message: "Invalid value, should be integer docker image tag version between dots"}, err)

	PrepareDummyFile(t, dummyFilePath, IncorrectNetworkNameLengthYamlV03Fixture)
	variablesSchema.Location = dummyFilePath
	err = variablesSchema.Read()
	assert.Error(t, err)
	assert.Equal(t, ClientError{Message: "Network name is too long. Maximum allowed characters are 20"}, err)

	PrepareDummyFile(t, dummyFilePath, IncorrectNetworkNameStringYamlV03Fixture)
	variablesSchema.Location = dummyFilePath
	err = variablesSchema.Read()
	err = variablesSchema.ValidateSchemaVariables()
	assert.Error(t, err)
	assert.Equal(t, ClientError{Message: "Network name is invalid"}, err)

	PrepareDummyFile(t, dummyFilePath, IncorrectStringVariablesYamlV03Fixture)
	variablesSchema.Location = dummyFilePath
	err = variablesSchema.Read()
	assert.Error(t, err)
	assert.Equal(t, ClientError{Message: "Variable with key: region contains white space which is not allowed"}, err)

	PrepareDummyFile(t, dummyFilePath, IncorrectAwsRegionType)
	variablesSchema.Location = dummyFilePath
	err = variablesSchema.Read()
	assert.Error(t, err)
	expectedError := ClientError{
		Message: fmt.Sprintf(
			"kaz-uk-2 is not valid AWS region. Valid are: %s",
			ValidRegions,
		),
	}
	assert.Equal(t, expectedError, err)

	PrepareDummyFile(t, dummyFilePath, IncorrectAwsInstanceType)
	variablesSchema.Location = dummyFilePath
	err = variablesSchema.Read()
	assert.Error(t, err)
	expectedError = ClientError{
		Message: fmt.Sprintf(
			"%s is not valid %s, valid are: %s",
			"someInvalid-instance-type.large",
			"asg_instance_type",
			ValidInstanceTypes,
		),
	}
	assert.Equal(t, expectedError, err)

	PrepareDummyFile(t, dummyFilePath, IncorrectConsensusMechanismType)
	variablesSchema.Location = dummyFilePath
	err = variablesSchema.Read()
	assert.Error(t, err)
	expectedError = ClientError{
		Message: fmt.Sprintf(
			"Invalid %s, valid are: %s",
			"consensus_mechanism",
			ValidConsensusMechanisms,
		),
	}
	assert.Equal(t, expectedError, err)

	RemoveDummyFile(t, dummyFilePath)
}

func TestMapAwsInstanceMemroyAndCpu(t *testing.T) {
	variablesSchema := VariablesSchema{}
	dummyFilePath := "dummy.yml"
	PrepareDummyFile(t, dummyFilePath, YamlFixtureAwsInstanceTypeSet)
	variablesSchema.Location = dummyFilePath
	err := variablesSchema.Read()
	assert.Nil(t, err)
	err = variablesSchema.ValidateSchemaVariables()
	assert.Nil(t, err)
	assert.Equal(
		t,
		variablesSchema.Variables["ecs_memory"],
		ValidInstances["t2.xlarge"].Memory,
	)

	RemoveDummyFile(t, dummyFilePath)
}

func configureYaml(version float64, resourceType string) (fileBody string) {
	return fmt.Sprintf(
		YamlFixtureConfigurable,
		version,
		resourceType,
	)
}
