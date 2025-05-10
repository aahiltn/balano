package api

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"

	"github.com/watchakorn-18k/scalar-go"
)

func Start() {
	app := fiber.New()
	app.Use(logger.New())
	app.Use("/docs", func(c *fiber.Ctx) error {
		htmlContent, err := scalar.ApiReferenceHTML(&scalar.Options{
			SpecURL: "openapi.yaml",
			CustomOptions: scalar.CustomOptions{
				PageTitle: "Palaam API Documentation",
			},
			DarkMode: true,
		})

		if err != nil {
			return err
		}
		c.Type("html")
		return c.SendString(htmlContent)
	})
	app.Use(recover.New())
	app.Use(cors.New())
	app.Get("/hello", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(&fiber.Map{
			"message": "Hello World",
		})
	})
	log.Fatal(app.Listen(":3000"))
}
