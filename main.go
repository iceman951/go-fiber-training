package main

import (
	"calcal/config"
	"calcal/middleware"
	"calcal/pkg/components"
	"calcal/pkg/handlers"

	"github.com/baac-tech/zlogres"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/rs/zerolog"
	"github.com/spf13/viper"
	"ipanda.baac.tech/golib/cipherPayload"
)

var (
	corsEnable bool
	corsConfig cors.Config

	cipherPayloadEnable bool
	cipherPayloadConfig cipherPayload.Config

	requestIDContextKey string
)

func init() {
	config.Setup("./env")

	requestIDContextKey = middleware.InitializeRequestIDFieldName()
	corsEnable, corsConfig = middleware.InitializeCORS()
	cipherPayloadEnable, cipherPayloadConfig = middleware.InitializeCipherPayload()

	zerolog.TimeFieldFormat = config.TimeFormat
}

func main() {
	app := fiber.New(fiber.Config{
		// Prefork: true, // When set to true, this will spawn multiple Go processes listening on the same port [Default: false]
	})

	if corsEnable {
		app.Use(cors.New(corsConfig))
	}
	app.Use(recover.New())
	app.Use(requestid.New(requestid.Config{
		ContextKey: requestIDContextKey,
	}))
	app.Use(zlogres.New(zlogres.Config{
		RequestIDContextKey: requestIDContextKey,
	}))
	if cipherPayloadEnable {
		app.Use(cipherPayload.New(cipherPayloadConfig))
	}

	componentLoaded := components.Load()
	handlers.SetEndpoints(app, componentLoaded.Handlers)

	PORT := viper.GetString("Listening.Port")
	app.Listen(":" + PORT)
}
