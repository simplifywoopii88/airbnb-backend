package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/simplifywoopii88/airbnb-backend/database"
	"github.com/simplifywoopii88/airbnb-backend/routes"
	"github.com/simplifywoopii88/airbnb-backend/utils"
)

const port string = ":3000"

func main() {
	database.ConnectDB()
	app := fiber.New()
	routes.SetupRoutes(app)

	routes.SetupMiddleware(app)

	if err := app.Listen(port); err != nil {
		utils.HandleErr(err)
	}

}
