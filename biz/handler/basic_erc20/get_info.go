package basic_erc20_handler

import (
	"eth-kawa/biz/infra"
	"fmt"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

func (h *BasicErc20Handler) GetInfo() {
	// 查询符号
	symbol, err := infra.GlobalErc20.Symbol(&bind.CallOpts{})
	if err != nil {
		h.c.JSON(consts.StatusInternalServerError, utils.H{"message": err.Error()})
		return
	}
	fmt.Println("代币符号:", symbol) // 输出 "MBT"

	// 查询总供应量（返回的是最小单位，需除以10^18才是可读数量）
	totalSupply, err := infra.GlobalErc20.TotalSupply(&bind.CallOpts{})
	if err != nil {
		h.c.JSON(consts.StatusInternalServerError, utils.H{"message": err.Error()})
		return
	}
	fmt.Println("总供应量（最小单位）:", totalSupply.String())
	h.c.JSON(consts.StatusOK, utils.H{"symbol": symbol, "totalSupply": totalSupply.String()})
}
