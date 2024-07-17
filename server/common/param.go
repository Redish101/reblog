package common

import (
	"net/http"
	"reblog/internal/core"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

func Param(app *core.App, c fiber.Ctx, dest interface{}) (isVaild bool, resp error) {
	reqMethod := c.Method()

	if isVaild, resp = ValidateParams(app, c, dest); !isVaild {
		return false, resp
	}

	if err := c.Bind().Query(dest); err != nil {
		return false, RespFail(c, http.StatusBadRequest, "无效的参数", err)
	}

	if reqMethod == "POST" || reqMethod == "PUT" {
		if err := c.Bind().Form(dest); err != nil {
			return false, RespFail(c, http.StatusBadRequest, "无效的参数", err)
		}
	}

	return true, nil
}

func ValidateParams(app *core.App, c fiber.Ctx, dest interface{}) (isVaild bool, resp error) {
	validate := app.Validator()
	err := validate.Struct(dest)

	if err != nil {
		if ve, ok := err.(*validator.ValidationErrors); ok {
			return false, RespFail(c, http.StatusBadRequest, ve.Error(), nil)
		}
		return false, RespServerError(c, err)
	}

	return true, nil
}
