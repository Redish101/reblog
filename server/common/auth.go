package common

import (
	"net/http"
	"reblog/internal/auth"

	"github.com/gofiber/fiber/v3"
)

// Tips: 后期新增多身份管理将使用此函数
// func Auth(role string) func(c fiber.Ctx) error

// 身份认证
func Auth() func(c fiber.Ctx) error {
	return func(c fiber.Ctx) error {
		token := c.Get("Authorization")

		if !auth.ValidToken(token) {
			return RespFail(c, http.StatusUnauthorized, "token错误", nil)
		}

		return c.Next()
	}
}
