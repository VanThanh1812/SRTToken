package handle

import (
	"srttoken/models/data_response"
	"srttoken/models/data_body"
	"srttoken/models"
	"srttoken/models/data_error"
	"srttoken/connection"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"strings"
	"log"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

const key = `{"address":"3a120d99ac41d6e0630382752f9714d8772c5432","crypto":{"cipher":"aes-128-ctr","ciphertext":"83b3120a8480229514ea6910b5d2f747f5ff7180677ddcaa75c157fafe11abbf","cipherparams":{"iv":"c2a0f16234c69f540e8e68cd1596ab10"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"5fb96f34f3e2308533223766b54a3ea14c4d60ff3f138481068016a126d77ba6"},"mac":"e99288c6f5e7aea1ba37db0a8f2e5a362764d784cadac166cf3a480a3b589c14"},"id":"89804198-b175-463a-a050-95940f91c1d9","version":3}`
const number = 1000

func CheckLinkShare(Link string){

}

func RequestSendToken(BodyBuy data_body.BodyBuyToken) *data_response.ResponseBuyToken {

	SRTToken := connection.GetSRTToken()

	if SRTToken == nil {
		return &data_response.ResponseBuyToken{
			Data: nil,
			Err: &data_error.ErrorResponse{
				Message:"Token null",
				Code:404,
			},
		}
	}

	auth, err := bind.NewTransactor(strings.NewReader(key), BodyBuy.Password)
	if err != nil {
		log.Printf("Failed to create authorized transactor: %v", err)
		return &data_response.ResponseBuyToken{
			Data: nil,
			Err: &data_error.ErrorResponse{
				Message:"Failed password",
				Code:500,
			},
		}
	}

	tx, err := SRTToken.Transfer(auth, common.HexToAddress(BodyBuy.Address), big.NewInt(number))

	if err != nil {
		log.Printf("Failed to transactor: %v", err)
		return &data_response.ResponseBuyToken{
			Data: nil,
			Err: &data_error.ErrorResponse{
				Message:"Failed to transactor",
				Code:500,
			},
		}
	}

	return &data_response.ResponseBuyToken{
		Data: &models.Transaction{
			Value: tx.Value(),
			To: tx.To().Hex(),
			TxHash:tx.Hash().Hex(),
		},
		Err: data_error.ErrorSuccess(),
	}
}