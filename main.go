package main

import (
	_ "SRTToken/routers"

	"github.com/astaxie/beego"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"github.com/ethereum/go-ethereum/common"
	"fmt"
	"SRTToken/contract"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"strings"
	"math/big"
)

const key = `{"address":"3a120d99ac41d6e0630382752f9714d8772c5432","crypto":{"cipher":"aes-128-ctr","ciphertext":"83b3120a8480229514ea6910b5d2f747f5ff7180677ddcaa75c157fafe11abbf","cipherparams":{"iv":"c2a0f16234c69f540e8e68cd1596ab10"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"5fb96f34f3e2308533223766b54a3ea14c4d60ff3f138481068016a126d77ba6"},"mac":"e99288c6f5e7aea1ba37db0a8f2e5a362764d784cadac166cf3a480a3b589c14"},"id":"89804198-b175-463a-a050-95940f91c1d9","version":3}`

func main() {

	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()

	conn, err := ethclient.Dial("/home/vanthanhbk/.ethereum/rinkeby/geth.ipc")

	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	// Instantiate the contract and display its name
	srttoken, err := contract.NewSRTToken(common.HexToAddress("0xaff6cfa079c49ae23b11d668b30c316e067c7fe8"), conn)
	if err != nil {
		log.Fatalf("Failed to instantiate a Token contract: %v", err)
	}
	name, err := srttoken.Name(nil)
	if err != nil {
		log.Fatalf("Failed to retrieve token name: %v", err)
	}

	fmt.Println("Token name:", name)

	auth, err := bind.NewTransactor(strings.NewReader(key), "thanhpv1234")
	if err != nil {
		log.Fatalf("Failed to create authorized transactor: %v", err)
	}

	//tx, err := srttoken.Transfer(auth, common.HexToAddress("0x99c92fEe17087FCb40da724da3783282D66b82e7"), big.NewInt(1))
	//
	//if err != nil {
	//	log.Fatalf("Failed to request token transfer: %v", err)
	//}
	//fmt.Printf("Transfer pending: 0x%x\n", tx.Hash())

	//approve, err := srttoken.Approve(auth, common.HexToAddress("0xaff6cfa079c49ae23b11d668b30c316e067c7fe8"), big.NewInt(2000))
	//
	//if err != nil {
	//	log.Fatalf("Failed to request token transfer: %v", err)
	//}
	//fmt.Printf("Transfer approve pending: 0x%x\n", approve.Hash())

	//s rt1 0x3A120D99ac41D6e0630382752F9714D8772c5432
	tx2, err := srttoken.TransferFrom(auth, common.HexToAddress("0x3A120D99ac41D6e0630382752F9714D8772c5432"), common.HexToAddress("0x99c92fEe17087FCb40da724da3783282D66b82e7"), big.NewInt(1500))

	if err != nil {
		log.Fatalf("Failed to request token transfer: %v", err)
	}

	fmt.Printf("Transfer 2 pending: 0x%x\n", tx2.Hash())

}
