package contract

import (
	erc20 "eth-kawa/example/smart_contracts/basic_erc20"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func LoadErc20(addressHex string, client *ethclient.Client) (*erc20.BasicErc20, error) {
	address := common.HexToAddress(addressHex)
	instance, err := erc20.NewBasicErc20(address, client)
	if err != nil {
		return nil, err
	}

	fmt.Println("contract is loaded")
	return instance, nil
}
