package main

import (
	"log"

	"example.com/indexer-wallet-txns/database"
	"example.com/indexer-wallet-txns/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// create app
	app := fiber.New()
	// setup routes
	routes.SetupRoutes(app)
	// Initialize the db
	database.InitDatabase()
	// start the server
	app.Listen(":3000")

	sqlDB, err := database.BlockDB.DB()
	if err != nil {
		log.Fatalf("Failed to get underlying database connection: %v", err)
	}
	defer sqlDB.Close()
}
