package handler

import (
	"reblog/internal/model"
	"reblog/internal/query"
	"reblog/server/common"

	"github.com/gofiber/fiber/v3"
)

//	@Summary		更新用户信息
//	@Description	管理员更新用户信息
//	@Tags			站点管理
//	@Param			username	path		string	true	"用户名"
//	@Param			nickname	formData	string	true	"昵称"
//	@Param			email		formData	string	true	"邮箱"
//	@Param			password	formData	string	true	"密码"
//	@Success		200			{object}	common.Resp
//	@Failure		400			{object}	common.Resp
//	@Router			/user/{username} [put]
func AdminUserUpdate(router fiber.Router) {
	router.Put("/user/:username", func(c fiber.Ctx) error {
		u := query.User

		username := c.Params("username")
		nickname := c.FormValue("nickname")
		email := c.FormValue("email")
		password := c.FormValue("password")

		if common.IsEmpty(username, nickname, email, password) {
			return common.RespMissingParameters(c)
		}

		_, err := u.Where(u.Username.Eq(username)).Updates(model.User{
			Username: username,
			Nickname: nickname,
			Email:    email,
			Password: password,
		})

		if err != nil {
			return common.RespServerError(c, err)
		}

		return common.RespSuccess(c, "操作成功, 登录态将失效", nil)
	}, common.Auth())
}
