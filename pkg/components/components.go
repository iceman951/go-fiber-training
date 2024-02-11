package components

import (
	"calcal/pkg/handlers"
	handlerHealth "calcal/pkg/handlers/health"
)

type Components struct {
	handlers.Handlers
}

func Load() Components {
	HealthHandler := loadHealthHandler()

	handlers := handlers.Handlers{
		Health: HealthHandler,
	}

	return Components{
		handlers,
	}
}

// Handler Loader

func loadHealthHandler() handlerHealth.Handler {
	return handlerHealth.NewHandler()
}
