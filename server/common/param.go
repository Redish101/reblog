package common

import (
	"net/http"
	"reblog/internal/core"
	"reblog/internal/log"

	"github.com/gofiber/fiber/v3"
	"github.com/mcuadros/go-defaults"
)

func Param(app *core.App, c fiber.Ctx, dest interface{}) (isVaild bool, resp error) {
	reqMethod := c.Method()

	defaults.SetDefaults(dest)

	if err := c.Bind().Query(dest); err != nil {
		return false, RespFail(c, http.StatusBadRequest, "无效的参数", err)
	}

	if reqMethod == "POST" || reqMethod == "PUT" {
		if err := c.Bind().Body(dest); err != nil {
			return false, RespFail(c, http.StatusBadRequest, "无效的参数", err)
		}
	}

	if isVaild, resp = ValidateParams(app, c, dest); !isVaild {
		return false, resp
	}

	return true, nil
}

func ValidateParams(app *core.App, c fiber.Ctx, dest interface{}) (isVaild bool, resp error) {
	validate := app.Validator()
	err := validate.Struct(dest)

	if err != nil {
		log.Warnf("[VALIDATOR] %v", err)
		return false, RespFail(c, http.StatusBadRequest, err.Error(), nil)
	}

	return true, nil
}
