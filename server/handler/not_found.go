package handler

import (
	"net/http"

	"github.com/redish101/reblog/internal/core"
	"github.com/redish101/reblog/server/common"

	"github.com/gofiber/fiber/v3"
)

func NotFound(app *core.App, router fiber.Router) {
	router.All("/*", func(c fiber.Ctx) error {
		return common.RespFail(c, http.StatusNotFound, "未知的接口", nil)
	})
}
