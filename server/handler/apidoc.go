package handler

import (
	"reblog/internal/core"

	"github.com/gofiber/fiber/v3"
)

func Apidoc(app *core.App, router fiber.Router) {
	router.Static("/apidoc", "./apidoc")
}
