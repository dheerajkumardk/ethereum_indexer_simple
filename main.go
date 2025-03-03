package main

import (
	"log"

	"example.com/indexer-wallet-txns/database"
	"example.com/indexer-wallet-txns/listener"
	"example.com/indexer-wallet-txns/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// create app
	app := fiber.New()

	// Initialize the db
	database.InitDatabase()

	// start block listener in a separate goroutine
	go listener.SubscribeBlocks()

	// setup routes
	routes.SetupRoutes(app)

	// start the server
	log.Fatal(app.Listen(":3000"))

	sqlDB, err := database.BlockDB.DB()
	if err != nil {
		log.Fatalf("Failed to get underlying database connection: %v", err)
	}
	defer sqlDB.Close()
}
