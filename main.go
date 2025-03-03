package main

import (
	"fmt"

	"example.com/indexer-wallet-txns/database"
	"example.com/indexer-wallet-txns/routes"
	"github.com/gofiber/fiber"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)


func main() {
	// create app
	app := fiber.New()
	// setup routes
	routes.SetupRoutes(app)
	// Initialize the db
	database.InitDatabase()
	
}
