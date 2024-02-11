package health

import (
	"calcal/pkg/services/health"
	"calcal/pkg/utils/library"

	"github.com/baac-tech/zlogwrap"
	"github.com/gofiber/fiber/v2"
)

type HandlerHealthInterface interface {
	HealthHandler(*fiber.Ctx) error
}

type Handler struct{}

func NewHandler() Handler {
	return Handler{}
}

func (h Handler) HealthHandler(c *fiber.Ctx) error {
	logger := zlogwrap.New(zlogwrap.Config{
		ServiceName: "HealthAPI",
	})

	if c.Method() != fiber.MethodGet {
		logger.SetField("event", map[string]interface{}{
			"error": fiber.ErrMethodNotAllowed,
		}).Error()
		return fiber.ErrMethodNotAllowed
	}

	c.Accepts(fiber.MIMEApplicationJSONCharsetUTF8)

	healthService := health.New()
	healthInfo := healthService.Get()

	resp := HealthResponse{
		Name:      healthInfo.Name,
		Version:   healthInfo.Version,
		ENV:       healthInfo.ENV,
		Status:    healthInfo.Status,
		Timestamp: library.GetTimestamp(),
	}

	return c.Status(fiber.StatusOK).JSON(resp)
}
