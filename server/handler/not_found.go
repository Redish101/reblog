package handler

import (
	"net/http"
	"reblog/server/common"

	"github.com/gofiber/fiber/v3"
)

func NotFound(router fiber.Router) {
	router.All("/*", func(c fiber.Ctx) error {
		return common.RespFail(c, http.StatusNotFound, "未知的接口", nil)
	})
}