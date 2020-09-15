package routes

import (
	"test-fiber/handlers"

	"github.com/gofiber/fiber/v2"
)

// Products is a routes for manage the products
func Products(app *fiber.App) {

	api := app.Group("/products")

	api.Get("/", handlers.GetProducts)
	api.Get("/:id", handlers.GetProduct)
	api.Post("/", handlers.CreateProduct)
	api.Delete("/:id", handlers.DeleteProduct)
}
