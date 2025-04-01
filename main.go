package main

import (
	"fiber-user-api/internal/database"
	"fiber-user-api/internal/routes"
	"github.com/gofiber/fiber/v3"
	"log"
)

func welcome(c fiber.Ctx) error {
	return c.SendString("Welcome 👋 to the API!")
}

func setupRoutes(app *fiber.App) {
	app.Get("/api", welcome)
	app.Post("/api/users", routes.CreateUser)
	app.Get("/api/users", routes.GetUsers)
	app.Get("/api/users/:id", routes.GetUser)
	app.Put("/api/users/:id", routes.UpdateUser)
	app.Delete("/api/users/:id", routes.DeleteUser)

	app.Post("/api/products", routes.CreateProduct)
	app.Get("/api/products", routes.GetProducts)
	app.Get("/api/products/:id", routes.GetProduct)
	app.Put("/api/products/:id", routes.UpdateProduct)
	app.Delete("/api/products/:id", routes.DeleteProduct)
	// order endpoint
	app.Post("/api/orders", routes.CreateOrder)
	app.Get("/api/orders", routes.GetOrders)
	app.Get("/api/orders/:id", routes.GetOrder)
	app.Put("/api/orders/:id", routes.UpdateOrder)
	app.Delete("/api/orders/:id", routes.DeleteOrder)

}

func main() {

	database.ConnectDB()

	app := fiber.New()

	app.Get("/", func(c fiber.Ctx) error {
		return c.SendString("Fiber API is running")
	})

	setupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
