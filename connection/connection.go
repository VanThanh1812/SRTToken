package connection

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"srttoken/contract"
	"github.com/ethereum/go-ethereum/common"
)

func tryConnectChain() *ethclient.Client {
	conn, err := ethclient.Dial("/home/vanthanhbk/.ethereum/rinkeby/geth.ipc")
	if err != nil {
		log.Println("Failed to connect to the Ethereum client: %v", err)
		return nil
	}

	return conn
}

func GetSRTToken() *contract.SRTToken {
	conn := tryConnectChain()

	srttoken, err := contract.NewSRTToken(common.HexToAddress("0xaff6cfa079c49ae23b11d668b30c316e067c7fe8"), conn)
	if err != nil {
		log.Println("Failed to instantiate a Token contract: %v", err)
		return nil
	}

	return srttoken
}
