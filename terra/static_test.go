package terra

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"index/suffixarray"
	"regexp"
	"testing"
)

func Test_CheckIntegrityOfStatics(t *testing.T) {
	quorumBody := StaticQuorum
	client := Client{}
	err := client.DefaultClient()
	assert.Nil(t, err)
	compareVarsToUsages(t, quorumBody)
}

func compareVarsToUsages(t *testing.T, body string) {
	variablesUsagePattern := regexp.MustCompile(`"\${var.[a-zA-Z0-9]*}"`)
	variableDeclarationPattern := `variable "%s"`
	variableNamePattern := regexp.MustCompile(`\.[a-zA-Z0-9]*`)
	bytesArray := suffixarray.New([]byte(body))
	results := bytesArray.FindAllIndex(variablesUsagePattern, -1)
	foundStrings := make([]string, 0)

	// Find all strings that are like "${var.something} but not "${var.something }"
	// It omits all redeclaration
	for _, result := range results {
		foundString := body[result[0]:result[1]]
		foundStrings = append(foundStrings, foundString)

		foundVariableNameSlice := variableNamePattern.FindStringSubmatch(foundString)

		if len(foundVariableNameSlice) < 1 {
			t.Errorf("%s variable name not found in string", foundString)
		}

		// Omit dot in regex match
		foundVariableName := foundVariableNameSlice[0][1:]

		// Seek for variable declaration
		declarationsPattern := fmt.Sprintf(variableDeclarationPattern, foundVariableName)
		variableDeclarationRegex := regexp.MustCompile(declarationsPattern)
		variableDeclarations := variableDeclarationRegex.FindAllString(body, -1)

		if len(variableDeclarations) < 1 {
			t.Errorf("%s variable declaration not found", foundVariableName)
		}
	}
}
