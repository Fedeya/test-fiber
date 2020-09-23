package routes

import (
	"github.com/fedeya/test-fiber/controllers"

	"github.com/gofiber/fiber/v2"
)

// Products is a routes for manage the products
func Products(app *fiber.App) {

	api := app.Group("/products")

	api.Get("/", controllers.GetProducts)
	api.Get("/:id", controllers.GetProduct)
	api.Post("/", controllers.CreateProduct)
	api.Put("/:id", controllers.UpdateProduct)
	api.Delete("/:id", controllers.DeleteProduct)
}
