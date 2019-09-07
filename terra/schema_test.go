package terra

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"path"
	"testing"
)

const (
	YamlV1Fixture = `version: 0.1
resourceType: variables
variables: 
  simpleKey: variable
`
	YamlFixtureConfigurable = `version: %v
resourceType: %s
variables:
  simpleKey: variable
`

	YamlFixtureNoVariables = `version: 0.1
resourceType: variables
`
	YamlFixtureWithHexUtils = `version: 0.1
resourceType: variables
variables:
  simpleKey: variable
  region:                'us-east-1'     ## You can set region for deployment here
  default_region:        'us-west-1'     ## If key region is not present it is default region setter
  profile:               'default'       ## It chooses profile from your ~/.aws config. If not present, profile is "default"
  aws_access_key_id:     'dummyValue'    ## It overrides access key id env variable. If omitted system env is used
  aws_secret_access_key: 'dummyValue'    ## It overrides secret access key env variable. If omitted system env is used
  genesis_gas_limit:      25		     ## Used to set genesis gas limit
  genesis_timestamp:      38	         ## Used to set genesis timestamp
  genesis_difficulty:     12             ## Used to set genesis difficulty
  genesis_nonce:          0              ## Used to set genesis nonce
  consensus_mechanism:    "instanbul"    ## Used to set consensus mechanism supported values are raft/istanbul
`

	NoSuchFileOrDirectoryMsg = "open %s: no such file or directory"
	NotValidExtMsg           = "%s is not in supported file types. Valid are: [.yml .yaml]"
)

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
	assert.Equal(t, "yaml: line 1: did not find expected ',' or '}'", err.Error())
	RemoveDummyFile(t, dummyFilePath)

	// It fails on invalid file path
	dummyFilePath = "dummyInvalidPath.yml"
	variablesSchema.Location = dummyFilePath
	err = variablesSchema.Read()
	assert.Error(t, err)
	assert.Equal(
		t,
		fmt.Sprintf("open %s: no such file or directory", dummyFilePath),
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
	assert.Equal(t, "yaml: line 1: did not find expected ',' or '}'", err.Error())
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
		fmt.Sprintf("%v version is not supported. Current version: %v", version, CurrentVersion),
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
		fmt.Sprintf("%s is not in supported resource types. Valid are: %s", resource, SupportedResourceTypes),
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
		"No variables found",
		err.Error(),
	)
	RemoveDummyFile(t, dummyFilePath)
}

func TestVariablesSchema_Read(t *testing.T) {
	variablesSchema := VariablesSchema{}
	dummyFilePath := "dummy.yml"
	PrepareDummyFile(t, dummyFilePath, YamlV1Fixture)
	variablesSchema.Location = dummyFilePath
	err := variablesSchema.Read()
	assert.Nil(t, err)
	assert.Equal(t, variablesSchema.Version, float64(0.1))
	assert.Equal(t, variablesSchema.Type, "variables")
	assert.NotNil(t, variablesSchema.Variables)
	variables := variablesSchema.Variables
	assert.Equal(t, "variable", variables["simpleKey"])
	RemoveDummyFile(t, dummyFilePath)
}

func TestVariablesSchema_Read_WithHexUtil(t *testing.T) {
	variablesSchema := VariablesSchema{}
	dummyFilePath := "dummy.yml"
	PrepareDummyFile(t, dummyFilePath, YamlFixtureWithHexUtils)
	variablesSchema.Location = dummyFilePath
	err := variablesSchema.Read()
	assert.Nil(t, err)

	assert.Equal(t, 4, len(VariablesKeyToHex))
	assert.Equal(t, "0x19", variablesSchema.Variables[VariablesKeyToHex[0]])
	assert.Equal(t, "0x26", variablesSchema.Variables[VariablesKeyToHex[1]])
	assert.Equal(t, "0xc", variablesSchema.Variables[VariablesKeyToHex[2]])
	assert.Equal(t, "0x0", variablesSchema.Variables[VariablesKeyToHex[3]])

	RemoveDummyFile(t, dummyFilePath)
}

func configureYaml(version float64, resourceType string) (fileBody string) {
	return fmt.Sprintf(
		YamlFixtureConfigurable,
		version,
		resourceType,
	)
}
