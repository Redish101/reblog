package common

import (
	"net/http"

	"github.com/gofiber/fiber/v3"
)

func GetFormValue(c fiber.Ctx, key string) string {
	data := c.FormValue(key, "")

	if data == "" {
		RespFail(c, http.StatusBadRequest, "参数缺失", nil)
	}

	return data
}