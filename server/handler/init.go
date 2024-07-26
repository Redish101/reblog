package handler

import (
	"net/http"

	"github.com/redish101/reblog/server/common"

	"github.com/redish101/reblog/internal/core"
	"github.com/redish101/reblog/internal/hash"
	"github.com/redish101/reblog/internal/model"

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
//	@Accept			json
//	@Produce		json
//	@Param			initParams	body		InitParams	true	"初始化参数"
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
