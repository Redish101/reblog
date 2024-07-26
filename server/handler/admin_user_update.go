package handler

import (
	"github.com/redish101/reblog/server/common"

	"github.com/redish101/reblog/internal/core"

	"github.com/gofiber/fiber/v3"
)

type AdminUserUpdateParams struct {
	Username string `json:"username" validate:"required,alphanum,min=3,max=32"`
	Nickname string `json:"nickname" validate:"required,min=2,max=32"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6,max=32"`
}

//	@Summary		更新用户信息
//	@Description	管理员更新用户信息
//	@Tags			站点管理
//	@Param			adminUserUpdateParams	body		AdminUserUpdateParams	true	"用户信息"
//	@Success		200						{object}	common.Resp
//	@Failure		400						{object}	common.Resp
//	@Router			/user/{username} [put]
func AdminUserUpdate(app *core.App, router fiber.Router) {
	router.Put("/user/:username", func(c fiber.Ctx) error {
		u := app.Query().User

		var params AdminUserUpdateParams
		if isValid, resp := common.ValidateParams(app, c, &params); !isValid {
			return resp
		}

		_, err := u.Where(u.Username.Eq(params.Username)).Updates(params)

		if err != nil {
			return common.RespServerError(c, err)
		}

		return common.RespSuccess(c, "操作成功, 登录态将失效", nil)
	}, common.Auth(app))
}
