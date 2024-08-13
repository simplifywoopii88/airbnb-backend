package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/simplifywoopii88/airbnb-backend/database"
	"github.com/simplifywoopii88/airbnb-backend/utils"
)

func createRoom(c *fiber.Ctx) error {
	var room database.Room

	if err := c.BodyParser(&room); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	utils.PrintStruct(room)

	// validation

	result := database.DB.Create(&room)
	if result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": result.Error,
		})
	}

	return c.Status(fiber.StatusOK).JSON(room.Serialize())
}
