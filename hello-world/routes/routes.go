package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(app *fiber.App) {
	// User endpoint
	userRoute := app.Group("/api/users")
	userRoute.Post("/", CreateUser)
	userRoute.Get("/", GetUsers)
	userRoute.Get("/:id", GetUser)
	userRoute.Put("/:id", UpdateUser)
	userRoute.Delete(":id", DeleteUser)

	// Product endpoint
	productRoute := app.Group("/api/users")
	productRoute.Post("/", CreateProduct)
	productRoute.Get("/", GetProducts)
	productRoute.Get("/:id", GetProduct)
	productRoute.Put("/:id", UpdateProduct)
	productRoute.Put("/:id", DeleteProduct)

	// Order endpoint
	orderRoute := app.Group("/api/orders")
	orderRoute.Post("/", CreateOrder)
	orderRoute.Get("/", GetOrders)
	orderRoute.Get("/:id", GetOrder)
}

func SetupMiddleware(app *fiber.App) {
	app.Use(logger.New())
}
