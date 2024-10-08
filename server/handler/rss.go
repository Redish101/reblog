package handler

import (
	"github.com/redish101/reblog/internal/core"
	"github.com/redish101/reblog/internal/feed"
	"github.com/redish101/reblog/server/common"

	"github.com/gofiber/fiber/v3"
)

// @Summary		获取Feed
// @Description	获取包含所有文章的Feed
// @Tags			Feed
// @Produce		xml
// @Success		200	"RSS Feed"
// @Failure		500	{object}	common.Resp	"服务器错误"
// @Router			/rss [get]
func Feed(app *core.App, router fiber.Router) {
	router.Get("/feed", func(c fiber.Ctx) error {
		a := app.Query().Article

		limit := app.Config().Rss.Limit
		articles, err := a.Order(a.CreatedAt.Desc()).Where(a.Draft.Is(false)).Limit(limit).Find()

		if err != nil {
			return common.RespServerError(c, err)
		}

		rssString, err := feed.GenerateFeed(app, articles)

		if err != nil {
			return common.RespServerError(c, err)
		}

		c.Set("Content-Type", "application/xml")

		return c.SendString(rssString)
	})
}
