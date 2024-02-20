package common

import (
	"net/http"

	"github.com/gofiber/fiber/v3"
)

type Resp struct {
	Success bool        `json:"success"`
	Msg     string      `json:"msg"`
	Data    interface{} `json:"data"`
}

func RespSuccess(c fiber.Ctx, msg string, data interface{}) error {
	return c.JSON(Resp{
		Success: true,
		Msg:     msg,
		Data:    data,
	})
}

func RespSuccessWithStatus(c fiber.Ctx, code int, msg string, data interface{}) error {
	return c.Status(code).JSON(Resp{
		Success: true,
		Msg:     msg,
		Data:    data,
	})
}

func RespFail(c fiber.Ctx, code int, msg string, data interface{}) error {
	return c.Status(code).JSON(Resp{
		Success: false,
		Msg:     msg,
		Data:    data,
	})
}

func RespServerError(c fiber.Ctx, errs ...error) error {
	var msg string

	for _, err := range errs {
		msg += err.Error()
	}

	return RespFail(c, http.StatusInternalServerError, msg, nil)
}

func RespMissingParameters(c fiber.Ctx) error {
	return c.Status(http.StatusBadRequest).JSON(Resp{
		Success: false,
		Msg:     "参数缺失",
	})
}
