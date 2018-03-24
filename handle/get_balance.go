package handle

import (
	"srttoken/connection"
	"github.com/ethereum/go-ethereum/common"
)

func GetBalance(Address string) int64 {
	SRTToken := connection.GetSRTToken()
	if SRTToken == nil {
		return -2
	}
	balance, err := SRTToken.BalanceOf(nil, common.HexToAddress(Address))
	if err != nil {
		return -1
	}
	return balance.Int64()
}
