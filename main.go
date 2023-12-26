package main

import (
	"log"
	"test/database"
	"test/routes"

	"github.com/gofiber/fiber/v2"
)
func welcome(c *fiber.Ctx) error { 
	return c.SendString("Welcome to the Application")

}

func setupRoutes(app *fiber.App){
	// welcome endpoint
	app.Get("/api", welcome)
	// user endpoint
	app.Post("/api/users", routes.CreateUser)
	app.Get("api/users", routes.GetUsers)
	app.Get("api/users/:id", routes.GetUser)
	app.Put("api/users/:id", routes.UpdateUser)
	app.Delete("api/users/:id", routes.Deleteuser)
	// product endpoint
	app.Post("api/product", routes.CreateProduct)
	app.Get("api/product", routes.GetProducts)
	app.Get("api/product/:id", routes.GetProduct)
	app.Put("api/product/:id", routes.UpdateProduct)
	// order endpoint
	app.Post("api/order", routes.CreateOrder)
	app.Get("api/order", routes.GetOrders)
	app.Get("api/order/:id", routes.GetOrder)
}

func main() {
	database.ConnectDb()
	app := fiber.New()

	app.Get("/api", welcome)

	log.Fatal(app.Listen(":3000"))
}