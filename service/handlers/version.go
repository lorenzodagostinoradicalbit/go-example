package handlers

import (
	fiber "github.com/gofiber/fiber/v2"
	"github.com/lorenzodagostinoradicalbit/go-example/config"
	"github.com/lorenzodagostinoradicalbit/go-example/service/response"
	"github.com/spf13/viper"
)

func Version(c *fiber.Ctx) error {
	version := viper.GetString(config.VERSION)
	vObj := &response.VersionResponse{Version: version}
	return c.JSON(vObj)
}
