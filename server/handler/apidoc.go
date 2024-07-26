package handler

import (
	"github.com/redish101/reblog/internal/core"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/static"
)

func Apidoc(app *core.App, router fiber.Router) {
	router.Use(static.New("apidoc"))
}
