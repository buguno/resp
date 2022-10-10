package routes

import "github.com/gofiber/fiber/v2"

// NotFoundRoute func for describe 404 Error route.
func NotFoundRoute(app *fiber.App) {
	// Register new special route.
	app.Use(
		// Anonymous function.
		func(context *fiber.Ctx) error {
			// Return HTTP 404 status and JSON response.
			return context.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": true,
				"msg":   "sorry, endpoint is not found",
			})
		},
	)
}
