package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/itkmitl/workshop-registration/internal/config"
	"github.com/itkmitl/workshop-registration/internal/database"
	"github.com/itkmitl/workshop-registration/internal/routes"
)

func main() {

	cfg := config.Load()

	database.Connect(cfg)
	database.Migrate()
	database.Seed()

	app := fiber.New(fiber.Config{
		AppName:   "IT KMITL Workshop Registration API",
		BodyLimit: 10 * 1024 * 1024,
	})

	app.Use(recover.New())
	app.Use(logger.New(logger.Config{
		Format: "${time} | ${status} | ${latency} | ${method} ${path}\n",
	}))
	app.Use(cors.New(cors.Config{
		AllowOrigins:     cfg.CORSOrigin,
		AllowMethods:     "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders:     "Origin,Content-Type,Accept,Authorization",
		AllowCredentials: true,
	}))

	routes.Setup(app, cfg)

	log.Printf("🚀 Server starting on port %s", cfg.Port)
	if err := app.Listen(":" + cfg.Port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
