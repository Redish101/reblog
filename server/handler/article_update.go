package handler

import (
	"net/http"
	"reblog/internal/core"
	"reblog/internal/model"
	"reblog/server/common"

	"github.com/gofiber/fiber/v3"
)

// @Summary		更新文章
// @Description	根据slug更新文章的标题、描述和内容
// @Tags			文章
// @Param			slug	path		string		true	"文章的slug"
// @Param			title	formData	string		true	"文章的标题"
// @Param			desc	formData	string		true	"文章的描述"
// @Param			content	formData	string		true	"文章的内容"
// @Success		200		{object}	common.Resp	"更新成功"
// @Failure		400		{object}	common.Resp	"缺失参数"
// @Failure		404		{object}	common.Resp	"未知的文章"
// @Security		ApiKeyAuth
// @Router			/article/{slug} [put]
func ArticleUpdate(app *core.App, router fiber.Router) {
	router.Put("/:slug", func(c fiber.Ctx) error {
		a := app.Query().Article

		slug := c.Params("slug")

		title := c.FormValue("title")
		desc := c.FormValue("desc")
		content := c.FormValue("content")

		if common.IsEmpty(slug, title, desc, content) {
			return common.RespMissingParameters(c)
		}

		article, err := a.Where(a.Slug.Eq(slug)).First()

		if article == nil {
			return common.RespFail(c, http.StatusNotFound, "未知的文章", nil)
		}

		if err != nil {
			return common.RespServerError(c, err)
		}

		_, err = a.Where(a.Slug.Eq(slug)).Updates(model.Article{
			Title:   title,
			Desc:    desc,
			Content: content,
		})

		if err != nil {
			return common.RespServerError(c, err)
		}

		return common.RespSuccess(c, "更新成功", nil)
	}, common.Auth(app))
}
