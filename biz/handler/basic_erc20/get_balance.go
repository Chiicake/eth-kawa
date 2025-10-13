package basic_erc20_handler

import (
	"eth-kawa/biz/infra"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

func (h *BasicErc20Handler) GetBalance() {
	address := h.c.Query("address")
	balance, err := infra.GlobalErc20.BalanceOf(&bind.CallOpts{}, common.HexToAddress(address))
	if err != nil {
		h.c.JSON(consts.StatusInternalServerError, utils.H{"message": err.Error()})
		return
	}
	h.c.JSON(consts.StatusOK, utils.H{"balance": balance.String()})
}
