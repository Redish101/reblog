package handler

import (
	"net/http"
	"strconv"

	"github.com/redish101/reblog/internal/core"
	"github.com/redish101/reblog/server/common"

	"github.com/gofiber/fiber/v3"
)

type FriendUpdateParams struct {
	Name    string `json:"name"`
	Avatar  string `json:"avatar"`
	Desc    string `json:"desc"`
	URL     string `json:"url"`
	Visible bool   `json:"visible"`
}

//	@Summary		更新友情链接
//	@Description	更新友情链接
//	@Tags			友情链接
//	@Accept			json
//	@Produce		json
//	@Param			friendUpdateParams	body		FriendUpdateParams	true	"更新友情链接参数"
//	@Success		200					{object}	common.Resp{}		"更新友情链接成功"
//	@Failure		400					{object}	common.Resp			"请求参数错误"
//	@Failure		500					{object}	common.Resp			"服务器内部错误"
//	@Security		ApiKeyAuth
//	@Router			/friend/{id} [put]
func FriendUpdate(app *core.App, router fiber.Router) {
	router.Put("/:id", func(c fiber.Ctx) error {
		f := app.Query().Friend

		idStr := c.Params("id")

		id, err := strconv.Atoi(idStr)
		if err != nil {
			return common.RespFail(c, http.StatusBadRequest, "无效的ID", err)
		}

		var params FriendUpdateParams
		if isValid, resp := common.Param(app, c, &params); !isValid {
			return resp
		}

		if friend, err := f.Where(f.ID.Eq(uint(id))).First(); friend == nil {
			if err != nil {
				return common.RespServerError(c, err)
			}

			return common.RespFail(c, http.StatusNotFound, "链接不存在", err)
		} else {
			friend.Name = params.Name
			friend.Avatar = params.Avatar
			friend.Desc = params.Desc
			friend.URL = params.URL
			friend.Visible = params.Visible

			if err := f.Save(friend); err != nil {
				return common.RespServerError(c, err)
			}
		}

		return common.RespSuccess(c, "更新成功", nil)
	}, common.Auth(app))
}
