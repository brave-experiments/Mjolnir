package terra

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
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

	hex = ConvertInterfaceToHex("two")
	assert.Equal(t, "0x0", hex)

	hex = ConvertInterfaceToHex("dummy")
	assert.Equal(t, "0x0", hex)

	dummyValue := map[string]int{"key": intValue}
	hex = ConvertInterfaceToHex(dummyValue)
	assert.Equal(t, "0x0", hex)
}

func TestReadOutputLogVarFailure(t *testing.T) {
	TempDirPathLocation = ".apolloTest"
	fullFilePath := TempDirPathLocation + "/output.log"
	invalidKey := "hello"
	// It fails when no output dir
	err, foundKey := ReadOutputLogVar(invalidKey)
	assert.Error(t, err)
	assert.Equal(t, fmt.Sprintf("open %s: no such file or directory", fullFilePath), err.Error())
	assert.Equal(t, len(foundKey), 0)

	// It fails when no output file
	err = os.MkdirAll(TempDirPathLocation, 0777)
	assert.Nil(t, err)
	err, foundKey = ReadOutputLogVar(invalidKey)
	assert.Error(t, err)
	assert.Equal(t, fmt.Sprintf("open %s: no such file or directory", fullFilePath), err.Error())
	assert.Equal(t, len(foundKey), 0)

	// It fails when no key within file
	file := File{
		Location: fullFilePath,
		Body:     OutputAsAStringWithInvalidValues,
	}
	err = file.WriteFile()
	assert.Nil(t, err)
	err, foundKey = ReadOutputLogVar(invalidKey)
	assert.Error(t, err)
	assert.Equal(t, ClientError{fmt.Sprintf("%s not found in output", invalidKey)}, err)
	assert.Equal(t, len(foundKey), 0)

	// It fails when no value is present within found key
	//err, foundKey = ReadOutputLogVar("network_name")
	//assert.Error(t, err)
	//assert.Equal(t, ClientError{"Value not present"}, err)

	err = os.RemoveAll(TempDirPathLocation)
	assert.Nil(t, err)

	TempDirPathLocation = TempDirPath
}

func TestReadOutputLogVar(t *testing.T) {
	TempDirPathLocation = ".apolloTest"
	fullFilePath := TempDirPathLocation + "/output.log"
	err := os.MkdirAll(TempDirPathLocation, 0777)
	assert.Nil(t, err)
	file := File{
		Location: fullFilePath,
		Body:     OutputAsAStringWithoutHeaderFixture,
	}
	err = file.WriteFile()
	assert.Nil(t, err)
	keyToSeek := "bastion_host_ip"
	err, foundKey := ReadOutputLogVar(keyToSeek)
	assert.Nil(t, err)
	assert.Equal(t, "3.15.144.150", foundKey)
}
