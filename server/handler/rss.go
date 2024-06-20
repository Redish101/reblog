package handler

import (
	"reblog/internal/core"
	"reblog/internal/rss"
	"reblog/server/common"

	"github.com/gofiber/fiber/v3"
)

// @Summary		获取Rss
// @Description	获取包含所有文章的RSS
// @Tags			Rss
// @Produce		xml
// @Success		200	"RSS Feed"
// @Failure		500	{object}	common.Resp	"服务器错误"
// @Router			/rss [get]
func Rss(app *core.App, router fiber.Router) {
	router.Get("/rss", func(c fiber.Ctx) error {
		a := app.Query().Article

		articles, err := a.Find()

		if err != nil {
			return common.RespServerError(c, err)
		}

		rssString, err := rss.GenerateRSS(app, articles)

		if err != nil {
			return common.RespServerError(c, err)
		}

		c.Set("Content-Type", "application/xml")

		return c.SendString(rssString)
	})
}
