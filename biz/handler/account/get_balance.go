package account_handler

import (
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/ethereum/go-ethereum/common"
)

func (h *AccountHandler) GetBalance() {
	addressHex := h.c.Query("address")
	if addressHex == "" {
		h.c.JSON(consts.StatusBadRequest, utils.H{"message": "address is required"})
		return
	}
	address := common.HexToAddress(addressHex)
	balance, err := h.client.BalanceAt(h.ctx, address, nil)
	if err != nil {
		h.c.JSON(consts.StatusBadRequest, utils.H{"message": err.Error()})
		return
	}
	h.c.JSON(consts.StatusOK, utils.H{"balance": balance})
}
