package handler

import (
	"github.com/redish101/reblog/server/common"

	"github.com/redish101/reblog/internal/core"

	"github.com/gofiber/fiber/v3"
)

//	@Summary		获取站点信息
//	@Description	获取站点信息
//	@Tags			站点
//	@Success		200	{object}	common.Resp{data=model.Site}
//	@Router			/site [get]
func Site(app *core.App, router fiber.Router) {
	router.Get("/", func(c fiber.Ctx) error {
		s := app.Query().Site

		site, err := s.First()

		if err != nil {
			return common.RespServerError(c, err)
		}

		return common.RespSuccess(c, "操作成功", site)
	})
}
