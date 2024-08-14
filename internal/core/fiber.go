package core

import (
	"github.com/ChuqiCloud/acmeidc/internal/version"

	"github.com/gofiber/fiber/v3"
)

func GetFiber() *fiber.App {
	fiberApp := fiber.New(fiber.Config{
		AppName:      version.GetAppName(),
		ServerHeader: version.GetAppName(),
	})

	return fiberApp
}
