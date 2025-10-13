package main

import (
	"context"
	account_handler "eth-kawa/biz/handler/account"
	basic_erc20_handler "eth-kawa/biz/handler/basic_erc20"
	"eth-kawa/biz/infra"
	"eth-kawa/biz/subscribe"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"log"
	"math/big"
	"strconv"
)

func main() {
	infra.InitDB()
	infra.InitEthClient()
	infra.InitWSEthClient()
	infra.LoadErc20()

	go subscribe.Subscribe()
	h := server.Default()

	h.GET("/ping", func(ctx context.Context, c *app.RequestContext) {
		c.JSON(consts.StatusOK, utils.H{"message": "pong"})
	})

	// get block by block number
	h.GET("/block/data/:number", func(ctx context.Context, c *app.RequestContext) {
		client := infra.GlobalEthClient
		number := c.Param("number")
		blockNumber, err := strconv.ParseInt(number, 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		block, err := client.BlockByNumber(context.Background(), big.NewInt(blockNumber))
		if err != nil {
			log.Fatal(err)
		}

		data := make([]utils.H, 0, len(block.Transactions()))
		// for each tx
		for i, tx := range block.Transactions() {
			data = append(data, utils.H{
				"index": i,
				"hash":  tx.Hash().Hex(),
				"data":  string(tx.Data()),
			})
		}
		c.JSON(consts.StatusOK, utils.H{"message": "block", "data": data})
	})

	h.GET("/api/accounts", func(ctx context.Context, c *app.RequestContext) {
		account_handler.NewAccountHandler(infra.GlobalEthClient, ctx, c, infra.GlobalDB).GetBalance()
	})

	h.POST("/api/transfer", func(ctx context.Context, c *app.RequestContext) {
		account_handler.NewAccountHandler(infra.GlobalEthClient, ctx, c, infra.GlobalDB).Transfer()
	})

	h.GET("/api/erc20/info", func(ctx context.Context, c *app.RequestContext) {
		basic_erc20_handler.NewBasicErc20Handler(infra.GlobalEthClient, ctx, c, infra.GlobalDB).GetInfo()
	})

	h.POST("/api/erc20/transfer", func(ctx context.Context, c *app.RequestContext) {
		basic_erc20_handler.NewBasicErc20Handler(infra.GlobalEthClient, ctx, c, infra.GlobalDB).Transfer()
	})

	h.Spin()
}
