package handler

import (
	"net/http"
	"reblog/internal/model"
	"reblog/internal/query"
	"reblog/server/common"

	"github.com/gofiber/fiber/v3"
)

// @Summary		添加文章
// @Description	添加一篇新的文章
// @Tags			文章
// @Param			slug	path		string		true	"文章slug"
// @Param			title	formData	string		true	"文章标题"
// @Param			desc	formData	string		true	"文章描述"
// @Param			content	formData	string		true	"文章内容"
// @Success		200		{object}	common.Resp	"操作成功"
// @Failure		400		{object}	common.Resp	"缺少必要参数"
// @Failure		409		{object}	common.Resp	"slug已被其他文章使用"
// @Security		ApiKeyAuth
// @Router			/article/{slug} [post]
func ArticleAdd(router fiber.Router) {
	router.Post("/:slug", func(c fiber.Ctx) error {
		a := query.Article

		slug := c.Params("slug")

		title := c.FormValue("title")
		desc := c.FormValue("desc")
		content := c.FormValue("content")

		if common.IsEmpty(title, slug, desc, content) {
			return common.RespMissingParameters(c)
		}

		if slug == "list" {
			return common.RespFail(c, http.StatusForbidden, "不能将list作为slug", nil)
		}

		existingArticle, _ := a.Where(a.Slug.Eq(slug)).First()

		if existingArticle != nil && !existingArticle.DeletedAt.Valid {
			return common.RespFail(c, http.StatusConflict, "当前slug已被其他文章使用", nil)
		}

		article := &model.Article{
			Title:   title,
			Slug:    slug,
			Desc:    desc,
			Content: content,
		}

		err := a.Create(article)

		if err != nil {
			return common.RespServerError(c, err)
		}

		return common.RespSuccess(c, "操作成功", nil)
	}, common.Auth())
}
