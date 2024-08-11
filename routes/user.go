package routes

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/simplifywoopii88/airbnb-backend/database"
	"github.com/simplifywoopii88/airbnb-backend/models"
)

func createUser(c *fiber.Ctx) error {
	// user struct init
	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// validation
	// CheckEmail
	fmt.Println("TODO: Email Check 중복검사 기능")
	// fmt.Println("email: ", user.Email)

	// database 생성
	database.DB.Create(&user)

	serializedUser := user.Serialize()
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"result": serializedUser,
	})

}
