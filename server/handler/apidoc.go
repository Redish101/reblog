package handler

import "github.com/gofiber/fiber/v3"

func Apidoc(router fiber.Router) {
	router.Static("/apidoc", "./apidoc")
}