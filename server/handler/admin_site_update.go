package handler

import (
	"github.com/redish101/reblog/internal/core"
	"github.com/redish101/reblog/server/common"

	"github.com/gofiber/fiber/v3"
)

type AdminSiteUpdateParams struct {
	Name string `json:"name" validate:"required"`
	Url  string `json:"url" validate:"required,url"`
	Desc string `json:"desc" validate:"required"`
	Icon string `json:"icon"`
}

// @Summary		更新站点信息
// @Description	更新站点的名称、URL、描述和图标
// @Tags			站点管理
// @Param			adminSiteUpdateParams	body		AdminSiteUpdateParams	true	"站点信息"
// @Success		200						{object}	common.Resp				"操作成功, 部分主题可能需重新部署生效"
// @Failure		400						{object}	common.Resp				"缺少参数"
// @Security		ApiKeyAuth
// @Router			/admin/site [PUT]
func AdminSiteUpdate(app *core.App, router fiber.Router) {
	router.Put("/site", func(c fiber.Ctx) error {
		s := app.Query().Site

		var params AdminSiteUpdateParams
		if isValid, resp := common.Param(app, c, &params); !isValid {
			return resp
		}
		_, err := s.Where(s.ID.Eq(1)).Updates(params)

		if err != nil {
			return common.RespServerError(c, err)
		}

		return common.RespSuccess(c, "操作成功, 部分主题可能需重新部署生效", nil)
	}, common.Auth(app))
}
