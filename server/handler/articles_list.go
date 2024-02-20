package handler

import (
	"fmt"
	"net/http"
	"reblog/internal/model"
	"reblog/internal/query"
	"reblog/server/common"
	"strconv"

	"github.com/gofiber/fiber/v3"
)

type ArticlesListResp struct {
	Count    int64            `json:"count"`
	Articles []*model.Article `json:"articles"`
}

//	@Summary		分页获取文章列表
//	@Description	分页获取文章列表
//	@Tags			文章
//	@Accept			json
//	@Produce		json
//	@Param			pageIndex	query		int									false	"页码, 默认为1"
//	@Param			pageSize	query		int									false	"每页数量, 默认为10"
//	@Success		200			{object}	common.Resp{data=ArticlesListResp}	"成功返回文章列表"
//	@Failure		400			{object}	common.Resp							"参数不合法"
//	@Failure		500			{object}	common.Resp							"服务器内部错误"
//	@Router			/article/list [get]
func ArticlesList(router fiber.Router) {
	router.Get("/list", func(c fiber.Ctx) error {
		a := query.Article

		pageIndexStr := c.Query("pageIndex", "1")
		pageSizeStr := c.Query("pageSize", "10")

		if common.CheckEmpty(pageIndexStr, pageSizeStr) {
			return common.RespMissingParameters(c)
		}

		pageIndex, indexErr := strconv.Atoi(pageIndexStr)
		pageSize, sizeErr := strconv.Atoi(pageSizeStr)

		if indexErr != nil || sizeErr != nil {
			msg := fmt.Sprintf("参数不合法: %v %v", indexErr, sizeErr)

			return common.RespFail(c, http.StatusBadRequest, msg, nil)
		}

		articles, count, err := a.Order(a.CreatedAt.Desc()).FindByPage((pageIndex-1)*pageSize, pageSize)

		if err != nil {
			return common.RespServerError(c, err)
		}
		return common.RespSuccess(c, "操作成功", ArticlesListResp{
			Count:    count,
			Articles: articles,
		})
	})
}
