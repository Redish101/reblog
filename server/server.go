package server

import (
	"fmt"
	"io/fs"
	"log"
	v "reblog/config"
	"reblog/internal/auth"
	"reblog/internal/config"
	"reblog/internal/db"
	"reblog/internal/query"
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
	config.InitConfig()

	auth.SetKey()

	query.Use(db.DB())
	query.SetDefault(db.DB())

	app := fiber.New(fiber.Config{
		AppName:      v.GetAppName(),
		ServerHeader: v.GetAppName(),
	})

	api := app.Group("/api")

	uifs := ui.GetUIFS()

	// logger
	app.Use(logger.New())

	// cors
	app.Use(cors.New(cors.ConfigDefault))

	// apidoc
	h.Apidoc(app)

	// init
	h.Init(api)

	// admin
	admin := api.Group("/admin")

	h.AdminLogin(admin)
	h.AdminTokenState(admin)
	h.AdminSiteUpdate(admin)

	// article
	article := api.Group("/article")

	h.ArticleList(article)
	h.ArticleSlug(article)
	h.ArticleAdd(article)
	h.ArticleDelete(article)
	h.ArticleUpdate(article)

	// rss
	h.Rss(app)

	// site
	site := api.Group("/site")

	h.Site(site)

	// version
	h.Version(api)

	// dashboard
	if config.Config().Dashboard.Enable {
		dashboard(app, uifs)
	}

	// notFound
	h.NotFound(app)

	log.Fatalln(listen(app))
}

func listen(app *fiber.App) error {
	serverConfig := config.Config().Server

	port := serverConfig.Port
	prefork := serverConfig.Prefork

	listenConfig := fiber.ListenConfig{
		EnablePrefork: prefork,
	}

	listenUrl := fmt.Sprintf(":%d", port)

	return app.Listen(listenUrl, listenConfig)
}

func dashboard(app *fiber.App, uifs fs.FS) {
	// fiber无法直接获取到index.html并返回, WTF?
	app.Get("/", func(c fiber.Ctx) error {
		indexFile, err := uifs.Open("dist/index.html")

		if err != nil {
			panic(err)
		}

		return c.Type("html").SendStream(indexFile)
	})

	app.Use("/", filesystem.New(filesystem.Config{
		Root:       ui.GetUIFS(),
		PathPrefix: "dist",
	}))
}
