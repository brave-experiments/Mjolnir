package terra

import (
	"github.com/brave-experiments/apollo-devops/src/github.com/stretchr/testify/assert"
	"testing"
)

func TestConvertIntToHex(t *testing.T) {
	valueToTest := int64(2)
	hex := ConvertIntToHex(valueToTest)
	assert.Equal(t, "0x2", hex)
}
