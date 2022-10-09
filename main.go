package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()

	app.Get("/", func(context *fiber.Ctx) error {
		return context.JSON(&fiber.Map{"data": "Hello, World 👋!"})
	})

	app.Listen(":3000")
}
