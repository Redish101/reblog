package server

import (
	"io"

	h "github.com/ChuqiCloud/acmeidc/server/handler"

	"github.com/ChuqiCloud/acmeidc/internal/config"
	"github.com/ChuqiCloud/acmeidc/internal/core"
	"github.com/ChuqiCloud/acmeidc/internal/log"
	"github.com/ChuqiCloud/acmeidc/internal/plugin"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	fb_logger "github.com/gofiber/fiber/v3/middleware/logger"
	_ "gorm.io/gorm"
)

func Start() {
	log.Info("欢迎使用acmeidc")

	config := config.NewFromFile()
	app := core.NewApp(config)

	loadPlugins(app)
	app.Bootstrap()

	fb := app.Fiber()

	api := fb.Group("/api")

	// logger
	fb.Use(fb_logger.New(fb_logger.Config{
		Format: "[HTTP] ${time} | ${status} | ${latency} | ${ip} | ${method} | ${path}",
		Output: io.Discard,
		Done: func(c fiber.Ctx, logString []byte) {
			code := c.Response().StatusCode()

			if code >= 200 && code < 400 {
				if app.Dev() {
					log.Info(string(logString))
				}
			} else {
				log.Error(string(logString))
			}
		},
	}))

	// cors
	fb.Use(cors.New(cors.ConfigDefault))

	// version
	h.Version(app, api)

	// notFound
	h.NotFound(app, fb)

	log.Fatal(app.Listen())
}

func loadPlugins(app *core.App) {
	plugins := app.Config().Plugins

	for _, pluginPath := range plugins {
		plugin.LoadPlugin(app, pluginPath)
	}
}
