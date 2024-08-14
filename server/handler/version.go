package handler

import (
	"runtime"

	"github.com/ChuqiCloud/acmeidc/internal/core"
	"github.com/ChuqiCloud/acmeidc/internal/version"
	"github.com/ChuqiCloud/acmeidc/server/common"

	"github.com/gofiber/fiber/v3"
)

type RespVersion struct {
	AppName string `json:"appName"`
	Version string `json:"version"`
	Commit  string `json:"commit"`
	Runtime string `json:"runtime"`
}

func Version(app *core.App, router fiber.Router) {
	router.Get("/version", func(c fiber.Ctx) error {
		return common.RespSuccess(c, "操作成功", RespVersion{
			AppName: version.GetAppName(),
			Version: version.Version,
			Commit:  version.Commit,
			Runtime: runtime.Version(),
		})
	})
}
