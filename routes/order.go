package routes

import (
	"errors"
	"test/database"
	"test/models"
	"time"

	"github.com/gofiber/fiber/v2"
)

type Order struct {
	ID       uint    `json:"id"`
	User     User    `json:"user"`
	Product  Product `json:"product"`
	CreateAt time.Time `json:"order_date"`
}

func CreateResponseOrder(order models.Order, user User, product Product) Order{
	return Order{
		ID: order.ID,
		User: user,
		Product: product,
		CreateAt: order.CreateAt,
	}
}

func CreateOrder (c *fiber.Ctx) error {
	var order models.Order

	if err := c.BodyParser(&order); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	var user models.User
	if err := findUser(order.UserRefer, &user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	var product models.Product
	if err := findProduct(order.ProductRefer, &product); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.Database.Db.Create(&order)

	responseUser := CreateResponseUser(user)
	responseProduct := CreateResponseProduct(product)
	responseOrder := CreateResponseOrder(order, responseUser, responseProduct)

	return c.Status(200).JSON(responseOrder)
}

func GetOrders (c *fiber.Ctx) error {
	orders := []models.Order{}
	database.Database.Db.Find(&orders)
	responseOrders := []Order{}

	for _,order := range orders {
		var user models.User
		var product models.Product

		database.Database.Db.Find(&user, "id = ?", order.UserRefer)
		database.Database.Db.Find(&product, "id=?", order.ProductRefer)

		responseOrder := CreateResponseOrder(order,CreateResponseUser(user), CreateResponseProduct(product))
		responseOrders = append(responseOrders, responseOrder)
	}
	return c.Status(200).JSON(responseOrders)
}

func FindOrder (id int, order *models.Order) error {
	database.Database.Db.Find(&order, "id=?", id)
	if order.ID == 0 {
		return errors.New("Order does not exist")
	}
	return nil
} 

func GetOrder( c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	var order models.Order

	if err != nil {
		return c.Status(400).JSON("PLease ensure that id is an integer")
	}
	if err := FindOrder(id, &order); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	var user models.User
	var product models.Product

	database.Database.Db.First(&user, order.UserRefer)
	database.Database.Db.First(&product, order.ProductRefer)
	responseUser := CreateResponseUser(user)
	reposnseProduct := CreateResponseProduct(product)

	responseOrder := CreateResponseOrder(order,responseUser,reposnseProduct )

	return c.Status(200).JSON(responseOrder)

	

}