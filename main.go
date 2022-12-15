package main

import (
	"log"
	"github.com/gofiber/fiber/v2"
	"github.com/venkatgears/Gofiber-Gorm-CRUD-API/database"
	"github.com/venkatgears/Gofiber-Gorm-CRUD-API/routes"
)

func welcome(c *fiber.Ctx) error {
	return c.SendString("Simple CRUD API using Go-Fiber & GORM & sqlite ")
}

func setupRoutes(app *fiber.App) {
	app.Get("/api",welcome)

	// Usermodel endpoints 
	app.Post("/api/users",routes.CreateUser)
	app.Get("api/users",routes.GetUsers)
	app.Get("api/users/:id",routes.GetUser)
	app.Put("api/users/:id",routes.UpdateUser)
	app.Delete("api/users/:id",routes.DeleteUser)

	// product endpoints 
	app.Post("/api/products",routes.CreateProduct)
	app.Get("/api/products",routes.Getproducts)
	app.Get("/api/products/:id",routes.GetProduct)
	app.Put("/api/products/:id",routes.Updateproduct)
	app.Delete("/api/products/:id",routes.DeleteProduct)

	// order endpointa 
	app.Post("api/orders",routes.CreateOrder)
	app.Get("api/orders",routes.GetOrders)
	app.Get("api/orders/:id",routes.GetOrder)
}

func main() {
	database.ConnectDb()
	app := fiber.New()
	setupRoutes(app)
	log.Fatal(app.Listen(":3000"))

}
