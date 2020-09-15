package main

import (
	"log"

	"test-fiber/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	app := fiber.New(fiber.Config{
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			message := "Internal Server Error"

			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
				message = e.Message
			}

			err = ctx.Status(code).JSON(fiber.Map{
				"status":     "error",
				"message":    message,
				"statusCode": code,
			})

			if err != nil {
				return ctx.Status(500).JSON(fiber.Map{
					"status":     "error",
					"message":    "Internal Server Error",
					"statusCode": 500,
				})
			}

			return nil
		},
	})
	app.Use(logger.New())
	app.Use(recover.New())

	routes.Routes(app)

	log.Fatal(app.Listen(":3000"))
}
