package main

import (
	"fmt"

	"github.com/gofiber/fiber"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupRoutes(app *fiber.App) {
	app.Get("/block/:blockNumber")
	app.Get("/block/:blockHash")
	app.Get("/blocks")
	app.Post("/blocks")
}

func main() {
	app := fiber.New()

}
