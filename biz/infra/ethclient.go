package infra

import "github.com/ethereum/go-ethereum/ethclient"

var GlobalEthClient *ethclient.Client

const ethClientURL = "http://localhost:8545"

func InitEthClient() *ethclient.Client {
	client, err := ethclient.Dial(ethClientURL)
	if err != nil {
		panic(err)
	}
	return client
}
