package routes

import "github.com/gofiber/fiber/v2"

// Routes manage the all routes
func Routes(app *fiber.App) {
	Products(app)
}
