package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/simplifywoopii88/airbnb-backend/database"
	"github.com/simplifywoopii88/airbnb-backend/routes"
)

const (
	port string = ":4000"
)

func main() {
	database.ConnectDB()
	app := fiber.New()

	//route
	routes.SetupRoutes(app)

	//middleware
	routes.SetupMiddleware(app)

	err := app.Listen(port)
	if err != nil {
		log.Println("Cannot start Listen and Serve!")
		log.Fatal(err.Error())
	}

}
