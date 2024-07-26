package main

import (
	"github.com/redish101/reblog/server/common"

	"github.com/redish101/reblog/internal/core"
	"github.com/redish101/reblog/internal/log"

	"github.com/gofiber/fiber/v3"
)

var _ core.Service = (*HelloPlugin)(nil)

type HelloPlugin struct {
	app *core.App
}

func NewHelloPlugin(app *core.App) core.Service {
	return &HelloPlugin{app: app}
}

func (p *HelloPlugin) Start() error {
	log.Infof("[HelloPlugin] Start")
	p.app.Fiber().All("/api/hello", func(c fiber.Ctx) error {
		return common.RespSuccess(c, "Hello from plugin!", nil)
	})

	return nil
}

func (p *HelloPlugin) Stop() error {
	return nil
}
