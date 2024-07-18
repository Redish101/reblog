package handler

import (
	"reblog/internal/core"
	"reblog/internal/model"
	"reblog/server/common"

	"github.com/gofiber/fiber/v3"
)

type FriendAddParams struct {
	Name   string `json:"name" validate:"required,max=32"`
	Avatar string `json:"avatar" validate:"required,url"`
	URL    string `json:"url" validate:"required,url"`
	Desc   string `json:"desc" validate:"max=256"`
}

//	@Summary		添加友情链接
//	@Description	添加友情链接
//	@Tags			友情链接
//	@Accept			json
//	@Produce		json
//	@Param			name	formData	string							true	"名称"
//	@Param			avatar	formData	string							true	"图标URL"
//	@Param			url		formData	string							true	"URL"
//	@Param			desc	formData	string							false	"描述"
//	@Success		200		{object}	common.Resp{data=model.Friend}	"添加友情链接成功"
//	@Failure		400		{object}	common.Resp						"请求参数错误"
//	@Failure		500		{object}	common.Resp						"服务器内部错误"
//	@Security		ApiKeyAuth
//	@Router			/friend [post]
func FriendAdd(app *core.App, router fiber.Router) {
	router.Post("/", func(c fiber.Ctx) error {
		f := app.Query().Friend

		var params FriendAddParams
		if isValid, resp := common.Param(app, c, &params); !isValid {
			return resp
		}

		friend := &model.Friend{
			Name:    params.Name,
			Avatar:  params.Avatar,
			URL:     params.URL,
			Desc:    params.Desc,
			Visible: false,
		}

		err := f.Create(friend)
		if err != nil {
			return common.RespServerError(c, err)
		}

		return common.RespSuccess(c, "操作成功", friend)
	})
}
