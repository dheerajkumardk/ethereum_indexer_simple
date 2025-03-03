package routes

import (
	"example.com/indexer-wallet-txns/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/block/blockNumber/:blockNumber", handlers.GetBlockByBlockNumber)
	app.Get("/block/blockHash/:blockHash", handlers.GetBlockByBlockHash)
	app.Get("/blocks", handlers.GetAllBlocks)
	app.Post("/blocks", handlers.AddBlock)
}
