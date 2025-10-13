package infra

import (
	"context"
	"eth-kawa/biz/contract/basic_erc20"
)

var GlobalErc20 *basic_erc20.BasicErc20

const GlobalPrivateKey = "49cbac6737c761593f7c0b4aad8a25c9607b68740ca04a96ef9608c0025ff673"

func DeployErc20() {
	GlobalErc20 = basic_erc20.DeployErc20(GlobalEthClient, context.Background(), GlobalPrivateKey)
}

const GlobalErc20Address = "1ff00af1c15d143bd4ca9f7d74ed7791b5cb353f"

func LoadErc20() {
	instance, err := basic_erc20.LoadErc20(GlobalErc20Address, GlobalEthClient)
	if err != nil {
		DeployErc20()
		return
	}
	GlobalErc20 = instance
}
