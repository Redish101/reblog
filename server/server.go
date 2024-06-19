package server

import (
	"io/fs"
	"reblog/internal/core"
	"reblog/internal/log"
	"reblog/internal/ui"
	h "reblog/server/handler"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/filesystem"
	"github.com/gofiber/fiber/v3/middleware/logger"
	_ "gorm.io/gorm"
)

//	@Title						reblog api
//	@Version					1.0
//	@License.name				GPL-V3
//	@Host						localhost:3000
//	@BasePath					/api
//	@Produce					json
//	@SecurityDefinitions.apikey	ApiKeyAuth
//	@In							header
//	@Name						Authorization
func Start() {
	log.Info("欢迎使用reblog")

	app := core.NewApp()

	fb := app.Fiber()

	api := fb.Group("/api")

	uifs := ui.GetUIFS()

	// logger
	fb.Use(logger.New())

	// cors
	fb.Use(cors.New(cors.ConfigDefault))

	// apidoc
	h.Apidoc(app, fb)

	// init
	h.Init(app, api)

	// admin
	admin := api.Group("/admin")

	h.AdminLogin(app, admin)
	h.AdminTokenState(app, admin)
	h.AdminSiteUpdate(app, admin)
	h.AdminUserInfo(app, admin)

	// article
	article := api.Group("/article")

	h.ArticleList(app, article)
	h.ArticleSlug(app, article)
	h.ArticleAdd(app, article)
	h.ArticleDelete(app, article)
	h.ArticleUpdate(app, article)

	// rss
	h.Rss(app, api)

	// site
	site := api.Group("/site")

	h.Site(app, site)

	// version
	h.Version(app, api)

	// dashboard
	if app.Config().Dashboard.Enable {
		dashboard(fb, uifs)
	}

	// notFound
	h.NotFound(app, fb)

	log.Fatal(app.Listen())
}

func dashboard(fb *fiber.App, uifs fs.FS) {
	// fiber无法直接获取到index.html并返回, WTF?
	fb.Get("/", func(c fiber.Ctx) error {
		indexFile, err := uifs.Open("dist/index.html")

		if err != nil {
			panic(err)
		}

		return c.Type("html").SendStream(indexFile)
	})

	fb.Use("/", filesystem.New(filesystem.Config{
		Root:       ui.GetUIFS(),
		PathPrefix: "dist",
	}))
}
