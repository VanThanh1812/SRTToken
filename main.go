package main

import (
	_ "srttoken/routers"

	"github.com/astaxie/beego"
	"srttoken/connection"
	"fmt"
)

func main() {

	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	srt := connection.GetSRTToken()
	name, _ := srt.Name(nil)
	fmt.Printf("Token name: 0x%x\n", name)

	//beego.Run()

	/*name, err := srttoken.Name(nil)
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

	fmt.Printf("Transfer 2 pending: 0x%x\n", tx2.Hash())*/

}
