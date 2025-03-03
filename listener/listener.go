package listener

import (
	"context"
	"fmt"
	"log"
	"os"

	"example.com/indexer-wallet-txns/database"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
)

func SubscribeBlocks() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	client, err := ethclient.Dial(os.Getenv("ETH_WSS_URL"))
	if err != nil {
		log.Fatalf("Error connecting to websocket: %v", err)
	}
	// create channel to receive latest block headers
	headers := make(chan *types.Header)

	sub, err := client.SubscribeNewHead(context.Background(), headers)
	if err != nil {
		log.Fatalf("Error subscribing %v", err)
	}

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case header := <-headers:

			block, err := client.BlockByHash(context.Background(), header.Hash())
			if err != nil {
				log.Fatal(err)
			}

			// create data to insert into db
			newBlock := database.Block{
				BlockNumber:  block.Number().Int64(),
				BlockHash:    block.Hash().String(),
				NumberOfTxns: int64(len(block.Transactions())),
				Miner:        block.Coinbase().Hex(),
			}

			// insert into db
			db := database.BlockDB
			result := db.Create(&newBlock)
			if result.Error != nil {
				log.Printf("Failed to insert block %d: %v", block.Number(), result.Error)
			} else {
				fmt.Printf("Stored block %d\n", block.Number())
			}
		}
	}
}
