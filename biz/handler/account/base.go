package account_handler

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/ethereum/go-ethereum/ethclient"
	"gorm.io/gorm"
)

type AccountHandler struct {
	client *ethclient.Client
	ctx    context.Context
	c      *app.RequestContext
	tx     *gorm.DB
}

func NewAccountHandler(client *ethclient.Client, ctx context.Context, c *app.RequestContext, tx *gorm.DB) *AccountHandler {
	return &AccountHandler{client, ctx, c, tx}
}
