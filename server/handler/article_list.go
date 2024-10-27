package handler

import (
	"github.com/redish101/reblog/internal/core"
	"github.com/redish101/reblog/internal/model"
	"github.com/redish101/reblog/internal/query"
	"github.com/redish101/reblog/server/common"

	"github.com/gofiber/fiber/v3"
)

type ArticleListParams struct {
	PageIndex int `json:"pageIndex" default:"1" validate:"min=1"`
	PageSize  int `json:"pageSize" default:"10" validate:"min=1"`
	Content bool `json:"content" default:"false"`
}

type ArticlesListResp struct {
	Count    int64            `json:"count"`
	Articles []*model.Article `json:"articles"`
}

//	@Summary		分页获取文章列表
//	@Description	分页获取文章列表
//	@Tags			文章
//	@Param			pageIndex	query		int									false	"页码, 默认为1"
//	@Param			pageSize	query		int									false	"每页数量, 默认为10"
//	@Param			content		query		bool								false	"是否返回文章内容, 默认为false"
//	@Success		200			{object}	common.Resp{data=ArticlesListResp}	"成功返回文章列表"
//	@Failure		400			{object}	common.Resp							"参数不合法"
//	@Failure		500			{object}	common.Resp							"服务器内部错误"
//	@Router			/article/list [get]
func ArticleList(app *core.App, router fiber.Router) {
	router.Get("/list", func(c fiber.Ctx) error {
		a := app.Query().Article

		var params ArticleListParams
		if isValid, resp := common.Param(app, c, &params); !isValid {
			return resp
		}

		includeDrafts := common.IsLogined(app, c)

		var query query.IArticleDo

		if !params.Content {
			query = a.Select(
				a.ID,
				a.Slug,
				a.CreatedAt,
				a.UpdatedAt,
				a.DeletedAt,
				a.Title,
				a.Slug,
				a.Desc,
				a.Draft,
			).Order(a.CreatedAt.Desc())
		} else {
			query = a.Select(
				a.ID,
				a.Slug,
				a.CreatedAt,
				a.UpdatedAt,
				a.DeletedAt,
				a.Title,
				a.Slug,
				a.Desc,
				a.Content,
				a.Draft,
			).Order(a.CreatedAt.Desc())
		}

		if !includeDrafts {
			query = query.Where(a.Draft.Is(false))
		}

		articles, count, err := query.FindByPage((params.PageIndex-1)*params.PageSize, params.PageSize)

		if err != nil {
			return common.RespServerError(c, err)
		}

		return common.RespSuccess(c, "操作成功", ArticlesListResp{
			Count:    count,
			Articles: articles,
		})
	})
}
