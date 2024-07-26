package handler

import (
	"github.com/redish101/reblog/server/common"

	"github.com/redish101/reblog/internal/core"
	"github.com/redish101/reblog/internal/model"

	"github.com/gofiber/fiber/v3"
)

type FriendListParams struct {
	PageIndex int `json:"pageIndex" default:"1" validate:"min=1"`
	PageSize  int `json:"pageSize" default:"10" validate:"min=1"`
}

type FriendListResp struct {
	Count   int64           `json:"count"`
	Friends []*model.Friend `json:"friends"`
}

//	@Summary		获取友情链接
//	@Description	分页获取友情链接, 若token有效则返回所有友情链接, 否则只返回可见的友情链接
//	@Tags			友情链接
//	@Accept			json
//	@Produce		json
//	@Param			pageIndex	query		integer			false	"页码，默认为1"
//	@Param			pageSize	query		integer			false	"每页大小，默认为10"
//	@Success		200			{object}	FriendListResp	"获取友情链接成功"
//	@Failure		500			{object}	common.Resp		"服务器内部错误"
//	@Router			/friend/list [get]
func FriendList(app *core.App, router fiber.Router) {
	router.Get("/list", func(c fiber.Ctx) error {
		f := app.Query().Friend

		var params FriendListParams
		if isValid, resp := common.Param(app, c, &params); !isValid {
			return resp
		}

		token := c.Get("Authorization")

		auth, err := core.AppService[*core.AuthService](app)
		if err != nil {
			return common.RespServerError(c, err)
		}

		friends, count, err := f.Order(f.CreatedAt.Desc()).FindByPage((params.PageIndex-1)*params.PageSize, params.PageSize)
		if err != nil {
			return common.RespServerError(c, err)
		}

		if auth.VerifyToken(token) {
			resp := FriendListResp{
				Count:   count,
				Friends: friends,
			}

			return common.RespSuccess(c, "操作成功", resp)
		}

		var cookedFriends []*model.Friend

		for _, friend := range friends {
			if friend.Visible {
				cookedFriends = append(cookedFriends, friend)
			}
		}

		resp := FriendListResp{
			Count:   count,
			Friends: cookedFriends,
		}

		return common.RespSuccess(c, "操作成功", resp)
	})
}
