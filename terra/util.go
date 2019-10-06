package terra

import (
	"bufio"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
	"os"
	"strconv"
	"strings"
)

func ConvertIntToHex(value int64) (hexInt string) {
	bigInt := big.NewInt(value)
	hexInt = hexutil.EncodeBig(bigInt)

	return hexInt
}

func ConvertInterfaceToHex(variable interface{}) (hexInt string) {
	switch valueInterface := variable.(type) {
	case float64:
		return ConvertIntToHex(int64(variable.(float64)))
	case float32:
		return ConvertIntToHex(int64(variable.(float32)))
	case int64:
		return ConvertIntToHex(variable.(int64))
	case int32:
		return ConvertIntToHex(int64(variable.(int32)))
	case int16:
		return ConvertIntToHex(int64(variable.(int16)))
	case int8:
		return ConvertIntToHex(int64(variable.(int8)))
	case int:
		return ConvertIntToHex(int64(variable.(int)))
	case uint64:
		return ConvertIntToHex(int64(variable.(uint64)))
	case uint32:
		return ConvertIntToHex(int64(variable.(uint32)))
	case uint16:
		return ConvertIntToHex(int64(variable.(uint16)))
	case uint8:
		return ConvertIntToHex(int64(variable.(uint8)))
	case uint:
		return ConvertIntToHex(int64(variable.(uint)))
	case string:
		float64Value, err := strconv.ParseFloat(valueInterface, 64)

		if nil != err {
			return ConvertIntToHex(int64(0))
		}

		return ConvertIntToHex(int64(float64Value))
	default:
		return ConvertIntToHex(int64(0))
	}
}

func ReadOutputLogVar(keyToRead string) (err error, readKey string) {
	err, deployDirNameLocator := fetchDeployDir()

	if nil != err {
		return err, readKey
	}

	file, err := os.Open(deployDirNameLocator + "/output.log")

	if nil != err {
		return err, readKey
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		text := scanner.Text()
		// whole key plus 3 signs (two spaces and one equal sign)
		valueWithoutSpacesLocator := len(keyToRead) + 3

		if false == strings.Contains(text, keyToRead) {
			continue
		}

		if len(text) < valueWithoutSpacesLocator {
			return ClientError{"Value not present"}, readKey
		}

		readKey = text[valueWithoutSpacesLocator:]
	}

	if len(readKey) < 1 {
		err = ClientError{fmt.Sprintf("%s not found in output", keyToRead)}
	}

	return err, readKey
}

func ReadSshLocation() (err error, sshKey string) {
	err, deployDirNameLocator := fetchDeployDir()

	if nil != err {
		return err, sshKey
	}

	fileLocator := deployDirNameLocator + "/id_rsa"

	return nil, fileLocator
}

func fetchDeployDir() (err error, deployNameLocator string) {
	tempDir, err := os.Open(TempDirPathLocation)

	if nil != err {
		return err, deployNameLocator
	}

	dirNames, err := tempDir.Readdir(0)

	if len(dirNames) < 1 {
		return ClientError{fmt.Sprintf("%s dir is empty", tempDir.Name())},
			deployNameLocator
	}

	deployDir := dirNames[0]
	deployNameLocator = TempDirPathLocation + "/" + deployDir.Name()

	return err, deployNameLocator
}
