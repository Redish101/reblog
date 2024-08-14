package handler

import (
	"net/http"

	"github.com/ChuqiCloud/acmeidc/internal/core"
	"github.com/ChuqiCloud/acmeidc/server/common"

	"github.com/gofiber/fiber/v3"
)

func NotFound(app *core.App, router fiber.Router) {
	router.All("/*", func(c fiber.Ctx) error {
		return common.RespFail(c, http.StatusNotFound, "未知的接口", nil)
	})
}
