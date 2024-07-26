package handler

import (
	"net/http"

	"github.com/redish101/reblog/server/common"

	"github.com/redish101/reblog/internal/core"
	"github.com/redish101/reblog/internal/model"

	"github.com/gofiber/fiber/v3"
)

type ArticleUpdateParams struct {
	Slug    string `json:"slug" validate:"required"`
	Title   string `json:"title" validate:"required"`
	Desc    string `json:"desc" validate:"required"`
	Content string `json:"content" validate:"required"`
}

//	@Summary		更新文章
//	@Description	根据slug更新文章的标题、描述和内容
//	@Tags			文章
//	@Param			slug				path		string				true	"文章的slug"
//	@Param			articleUpdateParams	body		ArticleUpdateParams	true	"文章更新参数"
//	@Success		200					{object}	common.Resp			"更新成功"
//	@Failure		400					{object}	common.Resp			"缺失参数"
//	@Failure		404					{object}	common.Resp			"未知的文章"
//	@Security		ApiKeyAuth
//	@Router			/article/{slug} [put]
func ArticleUpdate(app *core.App, router fiber.Router) {
	router.Put("/:slug", func(c fiber.Ctx) error {
		a := app.Query().Article

		var params ArticleUpdateParams
		params.Slug = c.Params("slug")
		if isValid, resp := common.Param(app, c, &params); !isValid {
			return resp
		}

		article, err := a.Where(a.Slug.Eq(params.Slug)).First()

		if article == nil {
			return common.RespFail(c, http.StatusNotFound, "未知的文章", nil)
		}

		if err != nil {
			return common.RespServerError(c, err)
		}

		_, err = a.Where(a.Slug.Eq(params.Slug)).Updates(model.Article{
			Title:   params.Title,
			Desc:    params.Desc,
			Content: params.Content,
		})

		if err != nil {
			return common.RespServerError(c, err)
		}

		return common.RespSuccess(c, "更新成功", nil)
	}, common.Auth(app))
}
