package basic_erc20_handler

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/ethereum/go-ethereum/ethclient"
	"gorm.io/gorm"
)

type BasicErc20Handler struct {
	client *ethclient.Client
	ctx    context.Context
	c      *app.RequestContext
	tx     *gorm.DB
}

func NewBasicErc20Handler(client *ethclient.Client, ctx context.Context, c *app.RequestContext, tx *gorm.DB) *BasicErc20Handler {
	return &BasicErc20Handler{
		client: client,
		ctx:    ctx,
		c:      c,
		tx:     tx,
	}
}
