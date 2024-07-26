package handler

import (
	"net/http"

	"github.com/redish101/reblog/server/common"

	"github.com/redish101/reblog/internal/core"

	"github.com/gofiber/fiber/v3"
)

type AdminLoginParams struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type AdminLoginResp struct {
	Token string `json:"token"`
}

//	@Summary		登录
//	@Description	管理员使用用户名和密码进行登录，若登录成功，返回token
//	@Tags			站点管理
//	@Param			adminLoginParams	body		AdminLoginParams					true	"登录凭据"
//	@Success		200					{object}	common.Resp{data=AdminLoginResp}	"登录成功"
//	@Failure		400					{object}	common.Resp							"缺少必要参数"
//	@Failure		401					{object}	common.Resp							"用户名或密码错误"
//	@Router			/admin/login [post]
func AdminLogin(app *core.App, router fiber.Router) {
	router.Post("/login", func(c fiber.Ctx) error {
		var params AdminLoginParams
		if isValid, resp := common.Param(app, c, &params); !isValid {
			return resp
		}

		auth, err := core.AppService[*core.AuthService](app)

		if err != nil {
			return common.RespServerError(c, err)
		}

		token := auth.GetToken(params.Username, params.Password)

		if token == "" {
			return common.RespFail(c, http.StatusUnauthorized, "用户名或密码错误", nil)
		}

		return common.RespSuccess(c, "登录成功", AdminLoginResp{
			Token: token,
		})
	})
}
