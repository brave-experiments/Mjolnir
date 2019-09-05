package terra

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOutputRecords_ParseOutputsFromJson(t *testing.T) {
	// It fails on invalid fixture
	outputRecords := OutputRecords{}
	outputRecords.ParseOutputsFromJson(InvalidOutputFixture)
	assert.Equal(t, 0, len(outputRecords.Records))

	// It parses records
	outputRecords = OutputRecords{}
	outputRecords.ParseOutputsFromJson(ProperOutputFixture)
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

	// Should parse multiple values to string
	// It parses records
	outputRecords = OutputRecords{}
	outputRecords.ParseOutputsFromJson(MultipleValuesOutputFixture)
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

	// Should parse object
	outputRecords = OutputRecords{}
	outputRecords.ParseOutputsFromJson(MultipleValuesOutputFixture)
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
}
