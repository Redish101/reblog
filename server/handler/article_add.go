package handler

import (
	"net/http"

	"github.com/redish101/reblog/internal/core"
	"github.com/redish101/reblog/internal/model"
	"github.com/redish101/reblog/server/common"

	"github.com/gofiber/fiber/v3"
)

type ArticleAddParams struct {
	Slug    string `json:"slug" validate:"required"`
	Title   string `json:"title" validate:"required"`
	Desc    string `json:"desc" validate:"required"`
	Content string `json:"content" validate:"required"`
	Draft   bool   `json:"draft"`
}

// @Summary		添加文章
// @Description	添加一篇新的文章
// @Tags			文章
// @Param			slug				path		string				true	"文章slug"
// @Param			articleAddParams	body		ArticleAddParams	true	"文章参数"
// @Success		200					{object}	common.Resp			"操作成功"
// @Failure		400					{object}	common.Resp			"缺少必要参数"
// @Failure		409					{object}	common.Resp			"slug已被其他文章使用"
// @Security		ApiKeyAuth
// @Router			/article/{slug} [post]
func ArticleAdd(app *core.App, router fiber.Router) {
	router.Post("/:slug", func(c fiber.Ctx) error {
		a := app.Query().Article

		var params ArticleAddParams
		params.Slug = c.Params("slug")
		if isValid, resp := common.Param(app, c, &params); !isValid {
			return resp
		}

		if params.Slug == "list" {
			return common.RespFail(c, http.StatusForbidden, "不能将list作为slug", nil)
		}

		existingArticle, _ := a.Where(a.Slug.Eq(params.Slug)).First()

		if existingArticle != nil && !existingArticle.DeletedAt.Valid {
			return common.RespFail(c, http.StatusConflict, "当前slug已被其他文章使用", nil)
		}

		article := &model.Article{
			Title:   params.Title,
			Slug:    params.Slug,
			Desc:    params.Desc,
			Content: params.Content,
			Draft:   &params.Draft,
		}

		err := a.Create(article)

		if err != nil {
			return common.RespServerError(c, err)
		}

		return common.RespSuccess(c, "操作成功", nil)
	}, common.Auth(app))
}
