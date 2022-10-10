package main

import (
	"resp-api/pkg/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	app := fiber.New()

	app.Use(recover.New())
	app.Use(logger.New(logger.Config{
		Format: "${time} - status:${status} - ${method} - path:${path} - time:${latency} \n" +
			"Request body: ${body}\n" +
			"Response body: ${resBody}\n",
		TimeZone: "America/Fortaleza",
	}))

	routes.EspRoute(app)
	routes.NotFoundRoute(app) // Register route for 404 Error.

	app.Listen(":3000")
}
