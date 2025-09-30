package block_service

import "github.com/ethereum/go-ethereum/ethclient"

import (
	"context"
	"gorm.io/gorm"
)

type BlockService struct {
	client *ethclient.Client
	db     *gorm.DB
	ctx    context.Context
}

func NewBlockService(client *ethclient.Client, db *gorm.DB, ctx context.Context) *BlockService {
	return &BlockService{
		client: client,
		db:     db,
		ctx:    ctx,
	}
}
