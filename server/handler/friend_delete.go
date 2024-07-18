package handler

import (
	"net/http"
	"reblog/internal/core"
	"reblog/server/common"
	"strconv"

	"github.com/gofiber/fiber/v3"
)

//	@Summary		删除友情链接
//	@Description	根据ID删除一个已存在的友情链接
//	@Tags			友情链接
//	@Accept			json
//	@Produce		json
//	@Param			id	path		integer		true	"友情链接的ID"
//	@Success		200	{object}	common.Resp	"删除友情链接成功"
//	@Failure		400	{object}	common.Resp	"无效的ID格式"
//	@Failure		404	{object}	common.Resp	"链接不存在"
//	@Failure		500	{object}	common.Resp	"服务器内部错误"
//	@Security		ApiKeyAuth
//	@Router			/friend/{id} [delete]
func FriendDelete(app *core.App, router fiber.Router) {
	router.Delete("/:id", func(c fiber.Ctx) error {
		f := app.Query().Friend

		idStr := c.Params("id")

		id, err := strconv.Atoi(idStr)
		if err != nil {
			return common.RespFail(c, http.StatusBadRequest, "无效的ID", err)
		}

		if friend, err := f.Where(f.ID.Eq(uint(id))).First(); friend != nil {
			if _, err := f.Delete(friend); err != nil {
				return common.RespFail(c, http.StatusInternalServerError, "删除链接失败", err)
			}

			if err != nil {
				return common.RespServerError(c, err)
			}

			return common.RespSuccess(c, "删除链接成功", nil)
		}

		return common.RespFail(c, http.StatusNotFound, "链接不存在", nil)
	}, common.Auth(app))
}
