package routes

import (
	"fiber-user-api/internal/database"
	"fiber-user-api/internal/models"
	"github.com/gofiber/fiber/v3"
)

type Product struct {
	ID           uint   `json:"id"`
	Name         string `json:"name"`
	SerialNumber string `json:"serial_number"`
}

func CreateResponseProduct(productModel models.Product) Product {
	return Product{
		ID:           productModel.ID,
		Name:         productModel.Name,
		SerialNumber: productModel.SerialNumber,
	}
}

func CreateProduct(c fiber.Ctx) error {
	var product models.Product

	if err := c.Bind().Body(&product); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	database.Database.Db.Create(&product)
	responseProduct := CreateResponseProduct(product)
	return c.Status(fiber.StatusCreated).JSON(responseProduct)
}

func GetProducts(c fiber.Ctx) error {
	var products []models.Product

	database.Database.Db.Find(&products)

	var responseProducts []Product
	for _, product := range products {
		responseProducts = append(responseProducts, CreateResponseProduct(product))
	}

	return c.Status(fiber.StatusOK).JSON(responseProducts)
}

func findProductByID(id string) (models.Product, error) {
	var product models.Product
	err := database.Database.Db.Where("id = ?", id).First(&product).Error
	if err != nil {
		return product, err
	}
	return product, nil
}

func GetProduct(c fiber.Ctx) error {
	id := c.Params("id")
	product, err := findProductByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(err.Error())
	}
	responseProduct := CreateResponseProduct(product)
	return c.Status(fiber.StatusOK).JSON(responseProduct)
}

func UpdateProduct(c fiber.Ctx) error {
	id := c.Params("id")
	product, err := findProductByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(err.Error())
	}

	var updatedProduct models.Product
	if err := c.Bind().Body(&updatedProduct); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	product.Name = updatedProduct.Name
	product.SerialNumber = updatedProduct.SerialNumber

	database.Database.Db.Save(&product)

	responseProduct := CreateResponseProduct(product)
	return c.Status(fiber.StatusOK).JSON(responseProduct)
}

func DeleteProduct(c fiber.Ctx) error {
	id := c.Params("id")
	product, err := findProductByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(err.Error())
	}

	database.Database.Db.Delete(&product)

	return c.SendStatus(fiber.StatusNoContent)
}
