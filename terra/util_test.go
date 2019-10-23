package terra

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvertInterfaceToHexFailure(t *testing.T) {
	hex, err := ConvertInterfaceToHex("two")
	assert.Equal(t, "0x0", hex)
	assert.Error(t, err)

	hex, err = ConvertInterfaceToHex("dummy")
	assert.Equal(t, "0x0", hex)
	assert.Error(t, err)

	dummyValue := map[string]int{"key": 4}
	hex, err = ConvertInterfaceToHex(dummyValue)
	assert.Equal(t, "0x0", hex)
	assert.Error(t, err)

	hex, err = ConvertInterfaceToHex(string(4))
	assert.Equal(t, "0x0", hex)
	assert.Error(t, err)

	hex, err = ConvertInterfaceToHex("0xWsA")
	assert.Equal(t, "0x0", hex)
	assert.Error(t, err)
}

func TestConvertInterfaceToHex(t *testing.T) {
	intValue := 2
	hex, err := ConvertInterfaceToHex(float64(intValue))
	assert.Equal(t, "0x2", hex)
	assert.Nil(t, err)

	hex, err = ConvertInterfaceToHex(float32(intValue))
	assert.Equal(t, "0x2", hex)
	assert.Nil(t, err)

	hex, err = ConvertInterfaceToHex(int64(intValue))
	assert.Equal(t, "0x2", hex)
	assert.Nil(t, err)

	hex, err = ConvertInterfaceToHex(int32(intValue))
	assert.Equal(t, "0x2", hex)
	assert.Nil(t, err)

	hex, err = ConvertInterfaceToHex(int16(intValue))
	assert.Equal(t, "0x2", hex)
	assert.Nil(t, err)

	hex, err = ConvertInterfaceToHex(int8(intValue))
	assert.Equal(t, "0x2", hex)
	assert.Nil(t, err)

	hex, err = ConvertInterfaceToHex(intValue)
	assert.Equal(t, "0x2", hex)
	assert.Nil(t, err)

	hex, err = ConvertInterfaceToHex(uint64(intValue))
	assert.Equal(t, "0x2", hex)
	assert.Nil(t, err)

	hex, err = ConvertInterfaceToHex(uint32(intValue))
	assert.Equal(t, "0x2", hex)
	assert.Nil(t, err)

	hex, err = ConvertInterfaceToHex(uint16(intValue))
	assert.Equal(t, "0x2", hex)
	assert.Nil(t, err)

	hex, err = ConvertInterfaceToHex(uint8(intValue))
	assert.Equal(t, "0x2", hex)
	assert.Nil(t, err)

	hex, err = ConvertInterfaceToHex(uint(intValue))
	assert.Equal(t, "0x2", hex)
	assert.Nil(t, err)

	hex, err = ConvertInterfaceToHex("0xE0000000")
	assert.Equal(t, "0xE0000000", hex)
	assert.Nil(t, err)

	hex, err = ConvertInterfaceToHex("0xEfAAA")
	assert.Equal(t, "0xEfAAA", hex)
}

func TestReadOutputLogVarFailure(t *testing.T) {
	TempDirPathLocation = ".mjolnirTestOutput"
	err := os.RemoveAll(TempDirPathLocation)
	assert.Nil(t, err)
	deployName := "dummyDeploy"
	deployNameLocator := TempDirPathLocation + "/" + deployName
	fullFilePath := deployNameLocator + "/output.log"
	invalidKey := "hello"

	// It fails when no output dir
	err, foundKey := ReadOutputLogVar(invalidKey)
	assert.Error(t, err)
	assert.Equal(t, fmt.Sprintf("open %s: no such file or directory", TempDirPathLocation), err.Error())
	assert.Equal(t, len(foundKey), 0)

	// It fails when no directory present
	err = os.MkdirAll(TempDirPathLocation, 0777)
	assert.Nil(t, err)
	err, foundKey = ReadOutputLogVar(invalidKey)
	assert.Error(t, err)
	assert.Equal(t, fmt.Sprintf("%s dir is empty", TempDirPathLocation), err.Error())
	assert.Equal(t, len(foundKey), 0)

	// It fails when there is no output file
	err = os.MkdirAll(deployNameLocator, 0777)
	assert.Nil(t, err)
	err, foundKey = ReadOutputLogVar(invalidKey)
	assert.Equal(t, fmt.Sprintf("open %s: no such file or directory", fullFilePath), err.Error())

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

	//It fails when no value is present within found key
	err, foundKey = ReadOutputLogVar("bastion_host_ip")
	assert.Error(t, err)
	assert.Equal(t, ClientError{"Value not present"}, err)

	err = os.RemoveAll(TempDirPathLocation)
	assert.Nil(t, err)

	TempDirPathLocation = TempDirPath
}

func TestReadOutputLogVar(t *testing.T) {
	TempDirPathLocation = ".mjolnirTestOutputLog"
	err := os.RemoveAll(TempDirPathLocation)
	assert.Nil(t, err)
	deployName := "dummyDeploy"
	deployNameDirLocator := TempDirPathLocation + "/" + deployName
	fullFilePath := deployNameDirLocator + "/output.log"
	err = os.MkdirAll(deployNameDirLocator, 0777)
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
	assert.Equal(t, "invalid.ip.666", foundKey)
	err = os.RemoveAll(TempDirPathLocation)
	assert.Nil(t, err)

	TempDirPathLocation = TempDirPath
}
