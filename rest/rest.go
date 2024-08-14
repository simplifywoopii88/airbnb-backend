package rest

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/simplifywoopii88/airbnb-backend/database"
	"github.com/simplifywoopii88/airbnb-backend/routes"
	"log"
)

func Start(port string) {
	if port == "" {
		port = "4000"
	}
	port = fmt.Sprintf(":%s", port)

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
