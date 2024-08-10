package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/simplifywoopii88/airbnb-backend/handlers"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/", handlers.HelloWorld)
}
