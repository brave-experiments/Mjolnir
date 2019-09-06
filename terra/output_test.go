package terra

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOutputRecords_FromJsonAsString(t *testing.T) {
	// Should return empty string
	outputRecords := OutputRecords{}
	stringOutput := outputRecords.FromJsonAsString(InvalidOutputFixture, false)
	assert.Equal(t, 0, len(outputRecords.Records))
	assert.Equal(t, 0, len(stringOutput))
	stringOutput = outputRecords.FromJsonAsString(InvalidOutputFixture, true)
	assert.Equal(t, 0, len(outputRecords.Records))
	assert.Equal(t, 0, len(stringOutput))

	// Should return parsed string from proper output fixture
	outputRecords = OutputRecords{}
	stringOutput = outputRecords.FromJsonAsString(ProperOutputFixture, false)
	assert.Equal(t, 10, len(outputRecords.Records))
	firstOutputRecord := outputRecords.Records[0]
	assert.IsType(t, OutputRecord{}, firstOutputRecord)
	assert.Equal(t, "_status", firstOutputRecord.Name)
	assert.Equal(t, false, firstOutputRecord.Sensitive)
	assert.Equal(t, "string", firstOutputRecord.Type)
	assert.Equal(
		t,
		"Completed!\n\nQuorum Docker Image         = quorumengineering/quorum:latest\nPrivacy Engine Docker Image = quorumengineering/tessera:latest\nNumber of Quorum Nodes      = 0\nECS Task Revision           = 2\nCloudWatch Log Group        = /ecs/quorum/cocroaches-attack\n",
		firstOutputRecord.Value,
	)
	assert.Equal(t, OutputAsAStringWithoutHeaderFixture, stringOutput)
	stringOutput = outputRecords.FromJsonAsString(ProperOutputFixture, true)
	assert.Equal(
		t,
		fmt.Sprintf("%s%s", ColorizedOutputPrefix, OutputAsAStringWithoutHeaderFixture),
		stringOutput,
	)

	// Should parse multiple values to string
	outputRecords = OutputRecords{}
	stringOutput = outputRecords.FromJsonAsString(MultipleValuesOutputFixture, false)
	assert.Equal(t, 10, len(outputRecords.Records))
	secondOutputRecord := outputRecords.Records[1]
	assert.IsType(t, OutputRecord{}, secondOutputRecord)
	assert.Equal(t, "bastion_host_dns", secondOutputRecord.Name)
	assert.Equal(t, false, secondOutputRecord.Sensitive)
	assert.Equal(t, "array", secondOutputRecord.Type)
	assert.Equal(
		t,
		"[\"\", \"\"]",
		secondOutputRecord.Value,
	)
	assert.Equal(
		t,
		OutputAsAStringFromMultipleValueTypes,
		stringOutput,
	)

	// Should return parser object with header
	outputRecords = OutputRecords{}
	stringOutput = outputRecords.FromJsonAsString(MultipleValuesOutputFixture, true)
	assert.Equal(t, 10, len(outputRecords.Records))
	thirdOutputRecord := outputRecords.Records[2]
	assert.IsType(t, OutputRecord{}, thirdOutputRecord)
	assert.Equal(t, "bastion_host_ip", thirdOutputRecord.Name)
	assert.Equal(t, false, thirdOutputRecord.Sensitive)
	assert.Equal(t, "map", thirdOutputRecord.Type)
	assert.Equal(
		t,
		"{\"ip\": \"3.15.144.150\"}",
		thirdOutputRecord.Value,
	)
	assert.Equal(
		t,
		fmt.Sprintf("%s%s", ColorizedOutputPrefix, OutputAsAStringFromMultipleValueTypes),
		stringOutput,
	)
}
