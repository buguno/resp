package main

import (
	"resp-api/pkg/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	routes.EspRoute(app)
	routes.NotFoundRoute(app) // Register route for 404 Error.

	app.Get("/", func(context *fiber.Ctx) error {
		return context.JSON(&fiber.Map{"data": "Hello, World 👋!"})
	})

	app.Listen(":3000")
}
