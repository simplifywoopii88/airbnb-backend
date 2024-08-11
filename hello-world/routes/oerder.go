package routes

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/simplifywoopii88/airbnb-backend/database"
	"github.com/simplifywoopii88/airbnb-backend/models"
	"time"
)

type Order struct {
	ID        uint      `json:"id"`
	User      User      `json:"user"`
	Product   Product   `json:"product"`
	CreatedAt time.Time `json:"order_date"`
}

func CreateResponseOrder(order models.Order, user User, product Product) Order {
	return Order{
		ID:        order.ID,
		User:      user,
		Product:   product,
		CreatedAt: order.CreatedAt,
	}
}

func CreateOrder(c *fiber.Ctx) error {
	var order models.Order

	if err := c.BodyParser(&order); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	var user models.User
	if err := findUser(order.UserRefer, &user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	var product models.Product
	if err := findProduct(order.UserRefer, &product); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.Database.DB.Create(&order)
	responseUser := CreateResponseUser(user)
	responseProduct := CreateResponseProduct(product)
	responeOrder := CreateResponseOrder(order, responseUser, responseProduct)

	return c.Status(200).JSON(responeOrder)
}

func GetOrders(c *fiber.Ctx) error {
	orders := []models.Order{}
	database.Database.DB.Find(&orders)
	responseOrders := []Order{}

	for _, order := range orders {
		var user models.User
		var product models.Product
		database.Database.DB.Find(&user, "id = ?", order.UserRefer)
		database.Database.DB.Find(&product, "id = ?", order.ProductRefer)
		responseOrder := CreateResponseOrder(order, CreateResponseUser(user), CreateResponseProduct(product))
		responseOrders = append(responseOrders, responseOrder)
	}

	return c.Status(200).JSON(responseOrders)
}

func findOrder(id int, order *models.Order) error {
	database.Database.DB.Find(order, "id = ?", id)
	if order.ID == 0 {
		return errors.New("order does not exists")
	}
	return nil
}

func GetOrder(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	var order models.Order

	if err != nil {
		return c.Status(400).JSON("Please ensure that :id is an integer")
	}

	if err := findOrder(id, &order); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	var user models.User
	var product models.Product

	database.Database.DB.First(&user, order.UserRefer)
	database.Database.DB.First(&product, order.ProductRefer)
	responseUser := CreateResponseUser(user)
	responseProduct := CreateResponseProduct(product)

	responseOrder := CreateResponseOrder(order, responseUser, responseProduct)

	//return c.Status(200).JSON(responseOrder)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "OK!",
		"result":  responseOrder,
	})
}
