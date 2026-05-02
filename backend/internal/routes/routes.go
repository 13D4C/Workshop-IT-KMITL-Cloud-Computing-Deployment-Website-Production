package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/itkmitl/workshop-registration/internal/config"
	"github.com/itkmitl/workshop-registration/internal/handlers"
	"github.com/itkmitl/workshop-registration/internal/middleware"
)

func Setup(app *fiber.App, cfg *config.Config) {
	api := app.Group("/api")

	api.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":  "healthy",
			"service": "it-kmitl-workshop-api",
		})
	})

	authHandler := handlers.NewAuthHandler(cfg)
	auth := api.Group("/auth")
	auth.Post("/register", authHandler.Register)
	auth.Post("/login", authHandler.Login)
	auth.Post("/logout", authHandler.Logout)
	auth.Get("/me", middleware.JWTProtected(cfg), authHandler.Me)

	activityHandler := handlers.NewActivityHandler()
	activities := api.Group("/activities")
	activities.Get("/", activityHandler.GetAll)
	activities.Get("/:id", activityHandler.GetByID)

	regHandler := handlers.NewRegistrationHandler()
	registrations := api.Group("/registrations", middleware.JWTProtected(cfg))
	registrations.Post("/", regHandler.Register)
	registrations.Get("/me", regHandler.GetMyRegistrations)
	registrations.Delete("/:id", regHandler.Cancel)
	registrations.Get("/check/:activityId", regHandler.CheckRegistration)
}
