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

func GetUsers(c fiber.Ctx) error {
	var users []models.User

	database.Database.Db.Find(&users)

	var responseUsers []User
	for _, user := range users {
		responseUsers = append(responseUsers, CreateResponseUser(user))
	}

	return c.Status(fiber.StatusOK).JSON(responseUsers)
}

func findUserByID(id string) (models.User, error) {
	var user models.User

	err := database.Database.Db.Where("id = ?", id).First(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func GetUser(c fiber.Ctx) error {
	id := c.Params("id")
	var user models.User

	//if err := database.Database.Db.Where("id = ?", id).First(&user).Error; err != nil {
	//	return c.Status(fiber.StatusNotFound).JSON(err.Error())
	//}

	user, err := findUserByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(err.Error())
	}

	responseUser := CreateResponseUser(user)
	return c.Status(fiber.StatusOK).JSON(responseUser)
}

func UpdateUser(c fiber.Ctx) error {
	id := c.Params("id")
	var user models.User

	//if err := database.Database.Db.Where("id = ?", id).First(&user).Error; err != nil {
	//	return c.Status(fiber.StatusNotFound).JSON(err.Error())
	//}

	user, err := findUserByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(err.Error())
	}

	var updatedUser models.User
	if err := c.Bind().Body(&updatedUser); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	user.FirstName = updatedUser.FirstName
	user.LastName = updatedUser.LastName

	database.Database.Db.Save(&user)

	responseUser := CreateResponseUser(user)
	return c.Status(fiber.StatusOK).JSON(responseUser)
}

func DeleteUser(c fiber.Ctx) error {
	id := c.Params("id")
	var user models.User

	user, err := findUserByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(err.Error())
	}

	database.Database.Db.Delete(&user)

	return c.Status(fiber.StatusNoContent).JSON(nil)

}
