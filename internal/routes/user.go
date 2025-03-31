package routes

import (
	"fiber-user-api/internal/database"
	"fiber-user-api/internal/models"
	"github.com/gofiber/fiber/v3"
)

type User struct {
	// serializer
	ID        uint   `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func CreateResponseUser(userModel models.User) User {
	return User{
		ID:        userModel.ID,
		FirstName: userModel.FirstName,
		LastName:  userModel.LastName,
	}
}

func CreateUser(c fiber.Ctx) error {
	var user models.User

	if err := c.Bind().Body(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	database.Database.Db.Create(&user)

	responseUser := CreateResponseUser(user)
	return c.Status(fiber.StatusCreated).JSON(responseUser)
}
