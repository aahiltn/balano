package main

import (
	"context"
	"log"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"palaam/internal/config"
	"palaam/internal/service"

	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
	"github.com/sethvargo/go-envconfig"
	"github.com/watchakorn-18k/scalar-go"
)

func main() {
	// Load environment variables
	var config config.Config
	if err := envconfig.Process(context.Background(), &config); err != nil {
		log.Fatalln("Error processing .env file: ", err)
	}

	app := fiber.New(fiber.Config{
		AppName: config.Application.Name,
	})
	appService := service.NewServer(config)

	routes.AppRouter(app, appService)

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

	go func() {
		if err := app.Listen(":" + config.Application.Port); err != nil {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit

	slog.Info("Shutting down server")
	if err := app.Shutdown(); err != nil {
		slog.Error("failed to shutdown server", "error", err)
	}

	slog.Info("Server shutdown")
}
