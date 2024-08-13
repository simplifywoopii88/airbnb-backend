package routes

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/simplifywoopii88/airbnb-backend/database"
)

func createUser(c *fiber.Ctx) error {
	// user struct init
	var user database.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// validation
	// CheckEmail
	if err := user.CheckDuplicatedEmail(); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// database 생성
	database.DB.Create(&user)

	serializedUser := user.Serialize()
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"result": serializedUser,
	})

}

func getUsers(c *fiber.Ctx) error {
	var users []database.User
	var serializedUsers []database.UserSerializer

	database.DB.Find(&users)

	for _, user := range users {
		serializedUsers = append(serializedUsers, user.Serialize())
	}
	return c.Status(fiber.StatusOK).JSON(serializedUsers)
}

func getUser(c *fiber.Ctx) error {

	// id parsing
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": ":id is an intger",
		})
	}

	// DB searching
	var user database.User
	if err := user.FindUser(id); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(user.Serialize())
}

func updateUser(c *fiber.Ctx) error {
	// id parsing
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": ":id is an intger",
		})
	}

	var user database.User
	if err := user.FindUser(id); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// bodyparsing

	payload := make(map[string]interface{})

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	var ok bool
	for k, v := range payload {
		switch k {
		case "name":
			user.Name, ok = v.(*string)
			if !ok {
				name := fmt.Sprintf("%v", v)
				user.Name = &name
			}
		case "is_host":
			user.IsHost, _ = v.(bool)
		case "gender":
			user.Gender, ok = v.(*string)
			if !ok {
				gender := fmt.Sprintf("%v", v)
				user.Gender = &gender
			}
		case "language":
			user.Language, _ = v.(string)
		case "currency":
			user.Currency, _ = v.(string)
		}
	}

	// timezone setting
	loc, _ := time.LoadLocation("Asia/Seoul")
	time.Local = loc
	updatedTime := time.Now()
	user.UpdatedAt = updatedTime

	database.DB.Save(&user)
	return c.Status(fiber.StatusOK).JSON(user.Serialize())
}

func deleteUser(c *fiber.Ctx) error {
	// findUser
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": ":id is an intger",
		})
	}

	var user database.User
	if err := user.FindUser(id); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	database.DB.Delete(&user)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "successful user delete!",
	})
}
