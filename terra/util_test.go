package terra

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConvertIntToHex(t *testing.T) {
	valueToTest := int64(2)
	hex := ConvertIntToHex(valueToTest)
	assert.Equal(t, "0x2", hex)
}

func TestConvertInterfaceToHex(t *testing.T) {
	intValue := 2
	hex := ConvertInterfaceToHex(float64(intValue))
	assert.Equal(t, "0x2", hex)

	hex = ConvertInterfaceToHex(float32(intValue))
	assert.Equal(t, "0x2", hex)

	hex = ConvertInterfaceToHex(int64(intValue))
	assert.Equal(t, "0x2", hex)

	hex = ConvertInterfaceToHex(int32(intValue))
	assert.Equal(t, "0x2", hex)

	hex = ConvertInterfaceToHex(int16(intValue))
	assert.Equal(t, "0x2", hex)

	hex = ConvertInterfaceToHex(int8(intValue))
	assert.Equal(t, "0x2", hex)

	hex = ConvertInterfaceToHex(intValue)
	assert.Equal(t, "0x2", hex)

	hex = ConvertInterfaceToHex(uint64(intValue))
	assert.Equal(t, "0x2", hex)

	hex = ConvertInterfaceToHex(uint32(intValue))
	assert.Equal(t, "0x2", hex)

	hex = ConvertInterfaceToHex(uint16(intValue))
	assert.Equal(t, "0x2", hex)

	hex = ConvertInterfaceToHex(uint8(intValue))
	assert.Equal(t, "0x2", hex)

	hex = ConvertInterfaceToHex(uint(intValue))
	assert.Equal(t, "0x2", hex)

	hex = ConvertInterfaceToHex(string(intValue))
	assert.Equal(t, "0x0", hex)

	hex = ConvertInterfaceToHex("0xE0000000")
	assert.Equal(t, "0xE0000000", hex)

	hex = ConvertInterfaceToHex("two")
	assert.Equal(t, "0x0", hex)

	hex = ConvertInterfaceToHex("dummy")
	assert.Equal(t, "0x0", hex)

	dummyValue := map[string]int{"key": intValue}
	hex = ConvertInterfaceToHex(dummyValue)
	assert.Equal(t, "0x0", hex)
}
