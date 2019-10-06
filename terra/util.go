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
	file, err := os.Open(TempDirPathLocation + "/output.log")

	if nil != err {
		return err, readKey
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		text := scanner.Text()

		if strings.Contains(text, keyToRead) {
			readKey = text
		}
	}

	if len(readKey) > 0 {
		return err, readKey
	}

	err = ClientError{fmt.Sprintf("%s not found in outout", keyToRead)}

	return err, readKey
}
