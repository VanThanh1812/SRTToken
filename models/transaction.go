package models

import "math/big"

type Transaction struct {
	TxHash string `json:"tx_hash"`
	BlockHeight int64 `json:"block_height"`
	From string `json:"from"`
	To string `json:"to"`
	Value *big.Int `json:"value"`
}

