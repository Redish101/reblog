package handler

import (
	"net/http"
	"reblog/internal/core"
	"reblog/server/common"

	"github.com/gofiber/fiber/v3"
)

//	@Summary		获取文章详情
//	@Description	根据slug获取文章详情
//	@Tags			文章
//	@Param			slug	path		string							true	"文章slug"
//	@Success		200		{object}	common.Resp{data=model.Article}	"操作成功"
//	@Failure		400		{object}	common.Resp						"缺少必要参数"
//	@Failure		404		{object}	common.Resp						"未知的slug"
//	@Router			/article/{slug} [get]
func ArticleSlug(app *core.App, router fiber.Router) {
	router.Get("/:slug", func(c fiber.Ctx) error {
		a := app.Query().Article

		slug := c.Params("slug")

		article, err := a.Where(a.Slug.Eq(slug)).First()

		if article == nil {
			return common.RespFail(c, http.StatusNotFound, "未知的slug", nil)
		}

		if err != nil {
			return common.RespServerError(c, err)
		}

		return common.RespSuccess(c, "操作成功", article)
	})
}
