package handler

import (
	"net/http"
	"reblog/internal/core"
	"reblog/internal/hash"
	"reblog/internal/model"
	"reblog/server/common"

	"github.com/gofiber/fiber/v3"
)

func isInited(app *core.App) bool {
	site, _ := app.Query().Site.First()

	return site != nil
}

type InitParams struct {
	Username string `json:"username" validate:"required"`
	Nickname string `json:"nickname" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
	Name     string `json:"name" validate:"required"`
	Url      string `json:"url" validate:"required,url"`
	Desc     string `json:"desc"`
	Icon     string `json:"icon"`
}

//	@Summary		初始化站点
//	@Description	使用给定的参数初始化站点
//	@Tags			站点管理
//	@Param			username	formData	string		true	"用户名"
//	@Param			nickname	formData	string		true	"昵称"
//	@Param			email		formData	string		true	"邮箱"
//	@Param			password	formData	string		true	"密码"
//	@Param			name		formData	string		true	"站点名称"
//	@Param			url			formData	string		true	"站点URL"
//	@Param			desc		formData	string		false	"站点描述"
//	@Param			icon		formData	string		false	"站点图标(base64格式)"
//	@Success		200			{object}	common.Resp	"初始化成功"
//	@Failure		400			{object}	common.Resp	"无效的邮箱或URL"
//	@Failure		403			{object}	common.Resp	"此站点已初始化"
//	@Router			/init [post]
func Init(app *core.App, router fiber.Router) {
	router.Post("/init", func(c fiber.Ctx) error {
		if isInited(app) {
			return common.RespFail(c, http.StatusForbidden, "此站点已初始化", nil)
		}

		var params InitParams
		if isValid, resp := common.Param(app, c, &params); !isValid {
			return resp
		}

		user := &model.User{
			Username: params.Username,
			Nickname: params.Nickname,
			Email:    params.Email,
			Password: hash.Hash(params.Password),
		}

		site := &model.Site{
			Name: params.Name,
			Url:  params.Url,
			Desc: params.Desc,
			Icon: params.Icon,
		}

		userErr := app.Query().User.Create(user)
		siteErr := app.Query().Site.Create(site)

		if userErr != nil || siteErr != nil {
			return common.RespServerError(c, userErr, siteErr)
		}

		return common.RespSuccess(c, "初始化成功", nil)
	})
}
