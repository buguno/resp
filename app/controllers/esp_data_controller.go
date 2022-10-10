package controllers

import (
	"resp-api/app/models"
	"resp-api/pkg/utils"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func CreateEspData(context *fiber.Ctx) error {
	// Create new EspData struct
	esp_data := &models.EspData{}

	// Check, if received JSON data is valid.
	if err := context.BodyParser(esp_data); err != nil {
		// Return status 400 and error message.
		return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	// Create database connection.
	// db, err := database.OpenDBConnection()
	// if err != nil {
	// 	// Return status 500 and database connection error.
	// 	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
	// 		"error": true,
	// 		"message":   err.Error(),
	// 	})
	// }

	// Set initialized default data for esp data:
	esp_data.Id = uuid.New()
	esp_data.CreatedAt = time.Now()
	esp_data.UpdatedAt = time.Now()

	// Create a new validator for a EspData model.
	errors := utils.ValidateEspDataStruct(*esp_data)
	if errors != nil {
		return context.Status(fiber.StatusBadRequest).JSON(errors)
	}

	// Return status 200 OK.
	return context.JSON(fiber.Map{
		"error":    false,
		"message":  nil,
		"esp_data": esp_data,
	})
}
