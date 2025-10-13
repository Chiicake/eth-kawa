package infra

import (
	"context"
	"eth-kawa/biz/contract"
	erc20 "eth-kawa/example/smart_contracts/basic_erc20"
)

var GlobalErc20 *erc20.BasicErc20

const GlobalPrivateKey = "869a589a1b188e9d1f84c74ad0dd6eaa6eb2f59460dd6c03bff0648736cd265e"

func DeployErc20() {
	GlobalErc20 = contract.DeployErc20(GlobalEthClient, context.Background(), GlobalPrivateKey)
}

const GlobalErc20Address = "0x5FbDB2315678afecb367f032d93F642f64180aa3"

func LoadErc20() {
	instance, err := contract.LoadErc20(GlobalErc20Address, GlobalEthClient)
	if err != nil {
		DeployErc20()
	}
	GlobalErc20 = instance
}
