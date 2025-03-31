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
