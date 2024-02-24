package handler

import (
	"reblog/internal/auth"
	"reblog/server/common"

	"github.com/gofiber/fiber/v3"
)

//	@Summary		获取token状态
//	@Description	获取当前token的状态
//	@Tags			站点管理
//	@Success		200	{object}	common.Resp{data=bool}	"若值为true则token有效"
//	@Security		ApiKeyAuth
//	@Router			/admin/tokenState [GET]
func AdminTokenState(router fiber.Router) {
	router.Get("/tokenState", func(c fiber.Ctx) error {
		token := c.Get("Authorization")
		state := auth.ValidToken(token)

		return common.RespSuccess(c, "操作成功", state)
	})
}
