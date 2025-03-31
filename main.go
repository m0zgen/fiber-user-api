package main

import (
	"github.com/gofiber/fiber/v3"
	"log"
)

func welcome(c fiber.Ctx) error {
	return c.SendString("Hello 👋 from API!")
}

func main() {
	app := fiber.New()
	app.Get("/api", welcome)

	app.Get("/", welcome)

	log.Fatal(app.Listen(":3000"))
}
