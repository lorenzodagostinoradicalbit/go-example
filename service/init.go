package service

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/lorenzodagostinoradicalbit/go-example/config"
	"github.com/lorenzodagostinoradicalbit/go-example/service/handlers"
	"github.com/spf13/viper"
)

func InitServer() {
	app := fiber.New()
	app.Use(logger.New())
	app.Use(cors.New())

	app.Get("/version", handlers.Version)

	app.Listen(fmt.Sprintf(":%s", viper.GetString(config.SERVER_PORT)))
}
