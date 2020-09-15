package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// Product is a simple type for products
type Product struct {
	ID       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	Price    float32   `json:"price"`
	Quantity int       `json:"quantity"`
}

var products []Product

func init() {
	id, _ := uuid.NewRandom()
	products = append(products, Product{
		ID:       id,
		Name:     "Laptop i7",
		Price:    2000,
		Quantity: 100,
	})
}

// GetProducts is a handler for return all products
func GetProducts(c *fiber.Ctx) error {
	return c.Status(200).JSON(products)
}

// GetProduct find product by id and return the product
func GetProduct(c *fiber.Ctx) error {
	id, _ := uuid.Parse(c.Params("id"))
	var product Product

	for i, p := range products {
		if p.ID == id {
			product = p
			break
		}
		if i == len(products)-1 {
			return fiber.NewError(fiber.StatusNotFound, "Product not Found")
		}
	}

	return c.Status(200).JSON(product)
}

// CreateProduct create a new product in the database
func CreateProduct(c *fiber.Ctx) error {
	product := new(Product)

	if err := c.BodyParser(product); err != nil {
		return fiber.ErrBadRequest
	}

	id, _ := uuid.NewRandom()

	*&product.ID = id

	go func() {
		products = append(products, *product)
	}()

	return c.JSON(product)
}

// DeleteProduct is a handler for delete product by id
func DeleteProduct(c *fiber.Ctx) error {

	id, _ := uuid.Parse(c.Params("id"))

	for i, product := range products {
		if product.ID == id {
			products = append(products[:i], products[i+1:]...)
			break
		}
		if i == len(products)-1 {
			return fiber.NewError(fiber.StatusNotFound, "Product not Found")
		}
	}

	return c.JSON(&fiber.Map{
		"message": "Product Deleted!",
	})
}
