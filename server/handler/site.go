package handler

import (
	"reblog/internal/query"
	"reblog/server/common"

	"github.com/gofiber/fiber/v3"
)

//	@Summary		获取站点信息
//	@Description	获取站点信息
//	@Tags			站点
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	common.Resp{data=model.Site}
//	@Router			/site [get]
func Site(router fiber.Router) {
	router.Get("/", func(c fiber.Ctx) error {
		s := query.Site

		site, err := s.First()

		if err != nil {
			return common.RespServerError(c, err)
		}

		return common.RespSuccess(c, "操作成功", site)
	})
}
