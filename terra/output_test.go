package terra

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOutputRecord_ParseOutputsFromJson(t *testing.T) {
	outputRecord := OutputRecord{}
	outputRecord.ParseOutputsFromJson(InvalidOutputFixture)
	assert.Equal(t, "", outputRecord.Name)
	assert.Equal(t, false, outputRecord.Sensitive)
	assert.Equal(t, "", outputRecord.Type)
	assert.Nil(t, outputRecord.Value)

	outputRecordName := "dummyRecord"
	outputRecord = OutputRecord{Name: outputRecordName}
	outputRecord.ParseOutputsFromJson(InvalidOutputFixture)
	assert.Equal(t, outputRecordName, outputRecord.Name)
	assert.Equal(t, false, outputRecord.Sensitive)
	assert.Equal(t, "", outputRecord.Type)
	assert.Nil(t, outputRecord.Value)

	outputRecordName = "bastion_host_ip"
	outputRecord = OutputRecord{Name: outputRecordName}
	outputRecord.ParseOutputsFromJson(ProperOutputFixture)
	assert.Equal(t, outputRecordName, outputRecord.Name)
	assert.Equal(t, false, outputRecord.Sensitive)
	assert.Equal(t, "", outputRecord.Type)
	assert.Nil(t, outputRecord.Value)
}
