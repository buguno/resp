package routes

import (
	"resp-api/app/controllers"

	"github.com/gofiber/fiber/v2"
)

func EspRoute(app *fiber.App) {
	//All routes related to esps comes here
	app.Post("/esp-data", controllers.CreateEspData)
}
