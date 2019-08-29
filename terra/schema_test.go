package terra

import (
	"fmt"
	"github.com/brave-experiments/apollo-devops/src/github.com/stretchr/testify/assert"
	"path"
	"testing"
)

const (
	YamlV1Fixture = `
apolloSchema: 0.1
resourceType: variables
  variables: 
    simpleKey: variable
`

	NoSuchFileOrDirectoryMsg = "open %s: no such file or directory"
	NotValidExtMsg           = "%s is not in supported types. Valid are: [.yml .yaml]"
)

func TestVariablesSchema_ReadFailure(t *testing.T) {
	variablesSchema := VariablesSchema{}
	err := variablesSchema.Read()
	assert.Error(t, err)
	assert.Equal(
		t,
		err.Error(),
		fmt.Sprintf(NoSuchFileOrDirectoryMsg, variablesSchema.Location),
	)

	variablesSchema.Location = "non-existing.yml"
	err = variablesSchema.Read()
	assert.Error(t, err)
	assert.Equal(
		t,
		err.Error(),
		fmt.Sprintf(NoSuchFileOrDirectoryMsg, variablesSchema.Location),
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
		err.Error(),
		fmt.Sprintf(NotValidExtMsg, path.Ext(dummyFilePath)),
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
}
