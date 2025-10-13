package main

import (
	erc20 "eth-kawa/example/smart_contracts/basic_erc20"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
)

func main() {
	client, err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		log.Fatal(err)
	}

	address := common.HexToAddress("1ff00af1c15d143bd4ca9f7d74ed7791b5cb353f")
	instance, err := erc20.NewBasicErc20(address, client)
	if err != nil {
		log.Fatal(err)
	}

	// 查询符号
	symbol, err := instance.Symbol(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("代币符号:", symbol)

	fmt.Println("contract is loaded")
	_ = instance
}
