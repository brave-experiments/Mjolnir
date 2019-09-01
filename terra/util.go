package terra

import (
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
)

func ConvertIntToHex(value int64) (hexInt string) {
	bigInt := big.NewInt(value)
	hexInt = hexutil.EncodeBig(bigInt)

	return hexInt
}
