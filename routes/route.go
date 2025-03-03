package routes

import "github.com/gofiber/fiber"

func SetupRoutes(app *fiber.App) {
	app.Get("/block/:blockNumber")
	app.Get("/block/:blockHash")
	app.Get("/blocks")
	app.Post("/blocks")
}
