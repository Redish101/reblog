package common

import (
	"net/http"

	"github.com/redish101/reblog/internal/core"

	"github.com/gofiber/fiber/v3"
)

// Tips: 后期新增多身份管理将使用此函数
// func Auth(role string) func(c fiber.Ctx) error

// 身份认证
func Auth(app *core.App) func(c fiber.Ctx) error {
	return func(c fiber.Ctx) error {
		token := c.Get("Authorization")

		auth, err := core.AppService[*core.AuthService](app)

		if err != nil {
			return RespServerError(c, err)
		}

		if !auth.VerifyToken(token) {
			return RespFail(c, http.StatusUnauthorized, "token错误", nil)
		}

		return c.Next()
	}
}

func IsLogined(app *core.App, c fiber.Ctx) bool {
	token := c.Get("Authorization")

	auth, err := core.AppService[*core.AuthService](app)

	if err != nil {
		return false
	}

	return auth.VerifyToken(token)
}
