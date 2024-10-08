package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(app *fiber.App) {
	// helloworld
	app.Get("/hello-world", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "hello world!",
		})
	})
}

func SetupMiddleware(app *fiber.App) {
	app.Use(logger.New())

	// user route
	userRoute := app.Group("/api/users")
	userRoute.Post("/", createUser)
	userRoute.Get("/", getUsers)
	userRoute.Get("/:id", getUser)
	userRoute.Patch("/:id", updateUser)
	userRoute.Delete("/:id", deleteUser)

	// rome route
	roomRoute := app.Group("/api/rooms")
	roomRoute.Post("/", createRoom)
}
