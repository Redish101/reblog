package handler

import (
	"reblog/internal/query"
	"reblog/server/common"

	"github.com/gofiber/fiber/v3"
)

type RespUserInfo struct {
	Username string `json:"username"`
	Nickname string `json:"nickname"`
}

//	@Summary		获取管理员信息
//	@Description	获取管理员信息
//	@Tags			站点管理
//	@Success		200	{object}	common.Resp{data=RespUserInfo}
//	@Router			/admin/userInfo [GET]
func AdminUserInfo(router fiber.Router) {
	router.Get("/userInfo", func(c fiber.Ctx) error {
		u := query.User

		user, err := u.First()

		if err != nil {
			return common.RespServerError(c, err)
		}

		return common.RespSuccess(c, "操作成功", RespUserInfo{
			Username: user.Username,
			Nickname: user.Nickname,
		})
	})
}