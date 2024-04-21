package server

import (
	"io/fs"
	"log"
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
//	@BasePath					/
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
		AppName:      "reblog",
		ServerHeader: "reblog",
	})

	uifs := ui.GetUIFS()

	// logger
	app.Use(logger.New())

	// cors
	app.Use(cors.New(cors.ConfigDefault))

	// dashboard
	dashboard(app, uifs)

	// apidoc
	h.Apidoc(app)

	// init
	h.Init(app)

	// admin
	admin := app.Group("/admin")

	h.AdminLogin(admin)
	h.AdminTokenState(admin)
	h.AdminSiteUpdate(admin)

	// article
	article := app.Group("/article")

	h.ArticleList(article)
	h.ArticleSlug(article)
	h.ArticleAdd(article)
	h.ArticleDelete(article)
	h.ArticleUpdate(article)

	// site
	site := app.Group("/site")

	h.Site(site)

	// notFound
	h.NotFound(app)

	log.Fatalln(app.Listen(":3000"))
}

func dashboard(app *fiber.App, uifs fs.FS) {
	app.Get("/", func(c fiber.Ctx) error {
		return c.Redirect().To("/dashboard/")
	})

	// fiber无法直接获取到index.html并返回, WTF?
	app.Get("/dashboard/", func(c fiber.Ctx) error {
		indexFile, err := uifs.Open("dist/index.html")

		if err != nil {
			panic(err)
		}

		return c.Type("html").SendStream(indexFile)
	})

	app.Use("/dashboard", filesystem.New(filesystem.Config{
		Root:       ui.GetUIFS(),
		PathPrefix: "dist",
	}))
}
