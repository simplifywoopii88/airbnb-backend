package routes

import (
	"fmt"

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

	fmt.Println(">>>>>>>>> room")
	utils.PrintStruct(room)

	return nil
}
