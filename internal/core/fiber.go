package core

import (
	"reblog/config"

	"github.com/gofiber/fiber/v3"
)

func GetFiber() *fiber.App {
	fiberApp := fiber.New(fiber.Config{
		AppName:      config.GetAppName(),
		ServerHeader: config.GetAppName(),
	})

	return fiberApp
}
