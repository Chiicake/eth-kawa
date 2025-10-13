package basic_erc20_handler

import (
	"crypto/ecdsa"
	"eth-kawa/biz/infra"
	"fmt"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"log"
	"math/big"
)

func (h *BasicErc20Handler) Transfer() {
	privateKeyHex := h.c.Query("privateKey")
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		h.c.JSON(consts.StatusBadRequest, utils.H{"message": err.Error()})
		return
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		h.c.JSON(consts.StatusBadRequest, utils.H{"message": "error casting public key"})
		return
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := h.client.PendingNonceAt(h.ctx, fromAddress)
	if err != nil {
		h.c.JSON(consts.StatusBadRequest, utils.H{"message": err.Error()})
		return
	}

	// 接收地址
	toAddressHex := h.c.Query("toAddress")
	toAddr := common.HexToAddress(toAddressHex)
	// 转账数量
	amount := h.c.Query("amount")
	amountBigInt, ok := new(big.Int).SetString(amount, 10)
	if !ok {
		h.c.JSON(consts.StatusBadRequest, utils.H{"message": "invalid amount"})
		return
	}

	// 构建交易签名器（同部署时的auth，需更新nonce）
	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.GasPrice, _ = h.client.SuggestGasPrice(h.ctx)

	// 调用transfer方法
	tx, err := infra.GlobalErc20.Transfer(auth, toAddr, amountBigInt)
	if err != nil {
		log.Fatal("转账失败:", err)
	}
	fmt.Println("转账交易哈希:", tx.Hash().Hex())
	h.c.JSON(consts.StatusOK, utils.H{"message": "transfer"})
}
