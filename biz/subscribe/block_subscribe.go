package subscribe

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
)

func Subscribe(url string) {
	client, err := ethclient.Dial(url)
	if err != nil {
		log.Fatal(err)
	}

	headers := make(chan *types.Header)
	sub, err := client.SubscribeNewHead(context.Background(), headers)
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case header := <-headers:
			fmt.Println("----------------------------------------")
			block, err := client.BlockByHash(context.Background(), header.Hash())
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println("block hash:", block.Hash().Hex())            // 0xbc10defa8dda384c96a17640d84de5578804945d347072e091b4e5f390ddea7f
			fmt.Println("block number:", block.Number().Uint64())     // 3477413
			fmt.Println("block time:", block.Time())                  // 1529525947
			fmt.Println("block nonce:", block.Nonce())                // 130524141876765836
			fmt.Println("block tx count:", len(block.Transactions())) // 7
			fmt.Println()
			// print each transactions
			for i, tx := range block.Transactions() {
				fmt.Println("tx index:", i)
				fmt.Println("tx Hash:", tx.Hash().Hex())
				from, err := types.Sender(types.NewEIP155Signer(tx.ChainId()), tx)
				if err != nil {
					log.Fatal(err)
				}
				fmt.Println("from:", from.Hex()) // 0x71C7656EC7ab88b098defB751B7401B5f6d8976F

				// to
				if tx.To() != nil {
					fmt.Println("to:", tx.To().Hex())
				}

				// data
				fmt.Println("data:", string(tx.Data()))
			}
			fmt.Println("----------------------------------------")
		}
	}
}
