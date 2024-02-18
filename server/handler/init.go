package handler

import (
	"net/http"
	"reblog/internal/hash"
	"reblog/internal/model"
	"reblog/internal/query"
	"reblog/server/common"
	"regexp"

	"github.com/gofiber/fiber/v3"
)

func isInited() bool {
	site, _ := query.Site.First()

	return site != nil
}

//	@Summary		初始化站点
//	@Description	使用给定的参数初始化站点
//	@Tags			站点管理
//	@Accept			json
//	@Produce		json
//	@Param			username	formData	string		true	"用户名"
//	@Param			nickname	formData	string		true	"昵称"
//	@Param			email		formData	string		true	"邮箱"
//	@Param			password	formData	string		true	"密码"
//	@Param			name		formData	string		true	"站点名称"
//	@Param			url			formData	string		true	"站点URL"
//	@Param			desc		formData	string		false	"站点描述"
//	@Param			icon		formData	string		false	"站点图标(base64格式)"
//	@Success		200			{object}	common.Resp	"初始化成功"
//	@Failure		400			{object}	common.Resp	"无效的邮箱或URL"
//	@Failure		403			{object}	common.Resp	"此站点已初始化"
//	@Router			/init [post]
func Init(router fiber.Router) {
	router.Post("/init", func(c fiber.Ctx) error {
		if isInited() {
			return common.RespFail(c, http.StatusForbidden, "此站点已初始化", nil)
		}

		username := c.FormValue("username")
		nickname := c.FormValue("nickname")
		email := c.FormValue("email")
		password := hash.Hash(c.FormValue("password"))
		name := c.FormValue("name")
		url := c.FormValue("url")
		desc := c.FormValue("desc")
		icon := c.FormValue("icon")

		if common.CheckEmpty(username, nickname, email, password, name, url) {
			return common.RespMissingParameters(c)
		}

		usernameRegex := `^[a-zA-Z0-9._%+-]`
		emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
		urlRegex := `^(https?:\/\/)?([\da-z\.-]+)\.([a-z\.]{2,6})([\/\w \.-]*)*\/?$`

		usernameMatched, _ := regexp.MatchString(usernameRegex, email)
		emailMatched, _ := regexp.MatchString(emailRegex, email)
		urlMatched, _ := regexp.MatchString(urlRegex, url)

		if !usernameMatched {
			return common.RespFail(c, http.StatusBadRequest, "无效的用户名", nil)
		}
		if !emailMatched {
			return common.RespFail(c, http.StatusBadRequest, "无效的电子邮件地址", nil)
		}
		if !urlMatched {
			return common.RespFail(c, http.StatusBadRequest, "无效的站点URL", nil)
		}

		user := &model.User{
			Username: username,
			Nickname: nickname,
			Email:    email,
			Password: password,
		}

		site := &model.Site{
			Name: name,
			Url:  url,
			Desc: desc,
			Icon: icon,
		}

		userErr := query.User.Create(user)
		siteErr := query.Site.Create(site)

		if userErr != nil || siteErr != nil {
			return common.RespServerError(c, userErr, siteErr)
		}

		return common.RespSuccess(c, "初始化成功", nil)
	})
}
