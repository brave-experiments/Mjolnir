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
	assert.Equal(t, 1, len(outputRecords.Records))
}
