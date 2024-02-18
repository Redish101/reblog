package server

import (
	"reblog/internal/auth"
	"reblog/internal/db"
	"reblog/internal/query"
	h "reblog/server/handler"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/logger"
	_ "gorm.io/gorm"
)

//	@Title			reblog api
//	@Version		1.0
//	@License.name	GPL-V3
//	@Host			localhost:3000
//	@BasePath		/
func Start() {
	auth.SetKey()

	query.Use(db.DB())
	query.SetDefault(db.DB())

	app := fiber.New(fiber.Config{
		AppName:      "reblog",
		ServerHeader: "reblog",
	})

	// logger
	app.Use(logger.New())

	// cors
	app.Use(cors.New(cors.ConfigDefault))

	// apidoc
	h.Apidoc(app)

	// init
	h.Init(app)

	// admin
	admin := app.Group("/admin")

	h.AdminLogin(admin)

	// articles
	articles := app.Group("/articles")

	h.ArticlesList(articles)
	h.ArticlesSlug(articles)
	h.ArticlesAdd(articles)

	// notFound
	h.NotFound(app)

	app.Listen(":3000")
}
