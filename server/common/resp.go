package common

import "github.com/gofiber/fiber/v3"

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
