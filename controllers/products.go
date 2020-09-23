package controllers

import (
	"github.com/fedeya/test-fiber/db"
	"github.com/fedeya/test-fiber/models"
	"github.com/gofiber/fiber/v2"
)

// GetProducts is a handler for return all products
func GetProducts(c *fiber.Ctx) error {
	var products []models.Product
	db := db.Get()

	db.Find(&products)

	return c.Status(200).JSON(products)
}

// GetProduct find product by id and return the product
func GetProduct(c *fiber.Ctx) error {
	var product models.Product

	db := db.Get()
	result := db.Find(&product, c.Params("id"))
	if result.RowsAffected == 0 {
		return fiber.ErrNotFound
	}

	return c.Status(200).JSON(product)
}

// CreateProduct create a new product in the database
func CreateProduct(c *fiber.Ctx) error {
	product := new(models.Product)
	db := db.Get()

	if err := c.BodyParser(product); err != nil {
		return fiber.ErrBadRequest
	}

	db.Create(product)

	return c.JSON(product)
}

// UpdateProduct is a handler for update product by id
func UpdateProduct(c *fiber.Ctx) error {
	var body map[string]interface{}
	var product models.Product
	db := db.Get()

	result := db.Find(&product, c.Params("id"))
	if result.RowsAffected == 0 {
		return fiber.ErrNotFound
	}

	if err := c.BodyParser(&body); err != nil {
		return fiber.ErrBadRequest
	}

	db.Model(&product).Updates(body)

	return c.JSON(product)
}

// DeleteProduct is a handler for delete product by id
func DeleteProduct(c *fiber.Ctx) error {

	db := db.Get()
	db.Delete(&models.Product{}, c.Params("id"))

	return c.JSON(&fiber.Map{
		"message": "Product Deleted!",
	})
}
