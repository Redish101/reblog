package handler

import (
	"net/http"
	"reblog/internal/auth"
	"reblog/server/common"

	"github.com/gofiber/fiber/v3"
)

type AdminLoginResp struct {
	Token string `json:"token"`
}

// @Summary		登录
// @Description	管理员使用用户名和密码进行登录，若登录成功，返回token
// @Tags			站点管理
// @Param			username	formData	string								true	"用户名或邮箱"
// @Param			password	formData	string								true	"密码"
// @Success		200			{object}	common.Resp{data=AdminLoginResp}	"登录成功"
// @Failure		400			{object}	common.Resp							"缺少必要参数"
// @Failure		401			{object}	common.Resp							"用户名或密码错误"
// @Router			/admin/login [post]
func AdminLogin(router fiber.Router) {
	router.Post("/login", func(c fiber.Ctx) error {
		username := c.FormValue("username")
		password := c.FormValue("password")

		if common.IsEmpty(username, password) {
			return common.RespMissingParameters(c)
		}

		token := auth.GetToken(username, password)

		if token == "" {
			return common.RespFail(c, http.StatusUnauthorized, "用户名或密码错误", nil)
		}

		return common.RespSuccess(c, "登录成功", AdminLoginResp{
			Token: token,
		})
	})
}
