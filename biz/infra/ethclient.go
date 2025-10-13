package infra

import "github.com/ethereum/go-ethereum/ethclient"

var GlobalEthClient *ethclient.Client
var GlobalWSEthClient *ethclient.Client

const ethClientURL = "http://localhost:8545"
const wsClientURL = "ws://localhost:8545"

func InitEthClient() {
	client, err := ethclient.Dial(ethClientURL)
	if err != nil {
		panic(err)
	}
	GlobalEthClient = client
}

func InitWSEthClient() {
	client, err := ethclient.Dial(wsClientURL)
	if err != nil {
		panic(err)
	}
	GlobalWSEthClient = client
}
