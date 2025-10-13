package basic_erc20

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func LoadErc20(addressHex string, client *ethclient.Client) (*BasicErc20, error) {
	address := common.HexToAddress(addressHex)
	instance, err := NewBasicErc20(address, client)
	if err != nil {
		return nil, err
	}

	return instance, nil
}
