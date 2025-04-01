package routes

import (
	"errors"
	"fiber-user-api/internal/database"
	"fiber-user-api/internal/models"
	"github.com/gofiber/fiber/v3"
	"time"
)

type Order struct {
	ID        int       `json:"id"`
	User      User      `json:"user"`
	Product   Product   `json:"product"`
	CreatedAt time.Time `json:"order_date"`
}

func CreateResponseOrder(orderModel models.Order, user User, product Product) Order {
	return Order{
		ID:        int(orderModel.ID),
		User:      user,
		Product:   product,
		CreatedAt: orderModel.CreatedAt,
	}
}

func CreateOrder(c fiber.Ctx) error {
	var order models.Order

	if err := c.Bind().Body(&order); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	var user models.User
	if err := database.Database.Db.Where("id = ?", order.UserRefer).First(&user).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(err.Error())
	}

	var product models.Product
	if err := database.Database.Db.Where("id = ?", order.ProductRefer).First(&product).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(err.Error())
	}

	database.Database.Db.Create(&order)

	responseUser := CreateResponseUser(user)
	responseProduct := CreateResponseProduct(product)

	responseOrder := CreateResponseOrder(order, responseUser, responseProduct)
	return c.Status(fiber.StatusCreated).JSON(responseOrder)
}

func GetOrders(c fiber.Ctx) error {
	var orders []models.Order

	database.Database.Db.Find(&orders)

	var responseOrders []Order
	for _, order := range orders {
		var user models.User
		if err := database.Database.Db.Where("id = ?", order.UserRefer).First(&user).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(err.Error())
		}

		var product models.Product
		if err := database.Database.Db.Where("id = ?", order.ProductRefer).First(&product).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(err.Error())
		}

		responseUser := CreateResponseUser(user)
		responseProduct := CreateResponseProduct(product)

		responseOrder := CreateResponseOrder(order, responseUser, responseProduct)
		responseOrders = append(responseOrders, responseOrder)
	}

	return c.Status(fiber.StatusOK).JSON(responseOrders)
}

func FindOrderByID(id string) (models.Order, error) {
	var order models.Order
	err := database.Database.Db.Where("id = ?", id).First(&order).Error
	if err != nil {
		return order, err
	}

	return order, nil
}

func GetOrder(c fiber.Ctx) error {
	id := c.Params("id")

	order, err := FindOrderByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(err.Error())
	}

	if order.ID == 0 {
		return c.Status(fiber.StatusNotFound).JSON(errors.New("order not found"))
	}

	var user models.User
	if err := database.Database.Db.Where("id = ?", order.UserRefer).First(&user).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(err.Error())
	}

	var product models.Product
	if err := database.Database.Db.Where("id = ?", order.ProductRefer).First(&product).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(err.Error())
	}

	responseUser := CreateResponseUser(user)
	responseProduct := CreateResponseProduct(product)

	responseOrder := CreateResponseOrder(order, responseUser, responseProduct)
	return c.Status(fiber.StatusOK).JSON(responseOrder)

}

func UpdateOrder(c fiber.Ctx) error {
	id := c.Params("id")
	order, err := FindOrderByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(err.Error())
	}

	if order.ID == 0 {
		return c.Status(fiber.StatusNotFound).JSON(errors.New("order not found"))
	}

	if err := c.Bind().Body(&order); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	database.Database.Db.Save(&order)

	return c.Status(fiber.StatusOK).JSON(order)
}

func DeleteOrder(c fiber.Ctx) error {
	id := c.Params("id")
	order, err := FindOrderByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(err.Error())
	}

	if order.ID == 0 {
		return c.Status(fiber.StatusNotFound).JSON(errors.New("order not found"))
	}

	database.Database.Db.Delete(&order)

	return c.SendStatus(fiber.StatusNoContent)
}
