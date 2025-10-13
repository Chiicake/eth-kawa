package account_handler

import (
	"crypto/ecdsa"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"math/big"
	"strconv"
)

func (h *AccountHandler) Transfer() {
	privateKeyHex := h.c.Query("private_key")
	if privateKeyHex == "" {
		h.c.JSON(consts.StatusBadRequest, utils.H{"message": "private_key is required"})
		return
	}

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

	toAddressHex := h.c.Query("to_address")
	if toAddressHex == "" {
		h.c.JSON(consts.StatusBadRequest, utils.H{"message": "to_address is required"})
		return
	}
	toAddress := common.HexToAddress(toAddressHex)

	value := h.c.Query("value")
	if value == "" {
		h.c.JSON(consts.StatusBadRequest, utils.H{"message": "value is required"})
		return
	}
	valueBigInt := big.NewInt(0)
	valueBigInt.SetString(value, 10)

	gasLimit := h.c.Query("gas_limit")
	if gasLimit == "" {
		h.c.JSON(consts.StatusBadRequest, utils.H{"message": "gas_limit is required"})
		return
	}
	var gasLimitBigUint64 uint64
	gasLimitBigUint64, err = strconv.ParseUint(gasLimit, 10, 64)
	if err != nil {
		h.c.JSON(consts.StatusBadRequest, utils.H{"message": err.Error()})
		return
	}

	gasPrice, err := h.client.SuggestGasPrice(h.ctx)
	if err != nil {
		h.c.JSON(consts.StatusBadRequest, utils.H{"message": err.Error()})
		return
	}

	var data []byte
	data = []byte(h.c.Query("data"))
	tx := types.NewTransaction(nonce, toAddress, valueBigInt, gasLimitBigUint64, gasPrice, data)
	// todo change to NewTx

	//chainID, err := h.client.NetworkID(h.ctx)
	//if err != nil {
	//	h.c.JSON(consts.StatusBadRequest, utils.H{"message": err.Error()})
	//	return
	//}
	//log.Info("chainID: %d", chainID)

	chainID := big.NewInt(1337)

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		h.c.JSON(consts.StatusBadRequest, utils.H{"message": err.Error()})
		return
	}

	err = h.client.SendTransaction(h.ctx, signedTx)
	if err != nil {
		h.c.JSON(consts.StatusBadRequest, utils.H{"message": err.Error()})
		return
	}

	h.c.JSON(consts.StatusOK, utils.H{"message": "tx sent", "tx_hash": signedTx.Hash().Hex()})

}
