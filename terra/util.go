package terra

import (
	"bufio"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var (
	ValidationHexPattern = `^(0x){1}[0-9a-fA-F]{0,128}$`
)

func ConvertInterfaceToHex(variable interface{}) (hexInt string, err error) {
	switch valueInterface := variable.(type) {
	case float64:
		return convertIntToHex(int64(variable.(float64))), err
	case float32:
		return convertIntToHex(int64(variable.(float32))), err
	case int64:
		return convertIntToHex(variable.(int64)), err
	case int32:
		return convertIntToHex(int64(variable.(int32))), err
	case int16:
		return convertIntToHex(int64(variable.(int16))), err
	case int8:
		return convertIntToHex(int64(variable.(int8))), err
	case int:
		return convertIntToHex(int64(variable.(int))), err
	case uint64:
		return convertIntToHex(int64(variable.(uint64))), err
	case uint32:
		return convertIntToHex(int64(variable.(uint32))), err
	case uint16:
		return convertIntToHex(int64(variable.(uint16))), err
	case uint8:
		return convertIntToHex(int64(variable.(uint8))), err
	case uint:
		return convertIntToHex(int64(variable.(uint))), err
	case string:
		matchesPattern, err := regexp.MatchString(ValidationHexPattern, valueInterface)

		if nil != err {
			return convertIntToHex(0), err
		}

		if matchesPattern {
			return string(valueInterface), err
		}

		float64Value, err := strconv.ParseInt(valueInterface, 10, 64)

		if nil != err {
			return convertIntToHex(0), err
		}

		return convertIntToHex(int64(float64Value)), err
	default:
		return convertIntToHex(int64(0)), ClientError{Message: "Invalid interface to hex conversion"}
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

func convertIntToHex(value int64) (hexInt string) {
	bigInt := big.NewInt(value)
	hexInt = hexutil.EncodeBig(bigInt)

	return hexInt
}

func isHexGreaterThanOrEqual(hex1 interface{}, hex2 interface{}) (isGreater bool) {
	hexString1, err := ConvertInterfaceToHex(hex1)
	hexString2, err := ConvertInterfaceToHex(hex2)

	if nil != err {
		return false
	}

	val1 := hexutil.MustDecodeUint64(hexString1)
	val2 := hexutil.MustDecodeUint64(hexString2)

	return val1 >= val2
}
