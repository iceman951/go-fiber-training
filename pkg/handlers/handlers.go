package handlers

import (
	handlerHealth "calcal/pkg/handlers/health"

	"github.com/gofiber/fiber/v2"
)

type Handlers struct {
	Health handlerHealth.HandlerHealthInterface
}

func SetEndpoints(app *fiber.App, h Handlers) {

	apiGroup := app.Group("/api")
	apiGroup.Get("/health", h.Health.HealthHandler)

	// v1 := apiGroup.Group(("/v1"))

	app.Use(func(c *fiber.Ctx) error { return c.SendStatus(fiber.StatusNotFound) })
}
