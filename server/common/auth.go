package common

import "github.com/gofiber/fiber/v3"

func Auth() func (c fiber.Ctx) error {
	return func(c fiber.Ctx) error {
		return c.SendString("aaaaaa")
	}
}