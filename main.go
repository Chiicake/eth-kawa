package main

import (
	"context"
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

const (
	wsURL     = "ws://localhost:8545"
	headerURL = "http://localhost:8545"
)

func main() {
	infra.InitDB()
	infra.InitEthClient()

	go subscribe.Subscribe(wsURL)
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

	h.Spin()
}
