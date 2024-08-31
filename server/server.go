package server

import (
	"io"
	"io/fs"
	"path/filepath"

	"github.com/redish101/reblog/server/common"
	h "github.com/redish101/reblog/server/handler"

	"github.com/redish101/reblog/internal/config"
	"github.com/redish101/reblog/internal/core"
	"github.com/redish101/reblog/internal/log"
	"github.com/redish101/reblog/internal/model"
	"github.com/redish101/reblog/internal/plugin"
	"github.com/redish101/reblog/internal/ui"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	fb_logger "github.com/gofiber/fiber/v3/middleware/logger"
	_ "gorm.io/gorm"
)

// @Title						reblog api
// @Version					1.0
// @License.name				GPL-V3
// @Host						localhost:3000
// @BasePath					/api
// @Produce					json
// @SecurityDefinitions.apikey	ApiKeyAuth
// @In							header
// @Name						Authorization
func Start() {
	log.Info("欢迎使用reblog")

	config := config.NewFromFile()
	app := core.NewApp(config)

	loadPlugins(app)
	app.Bootstrap()

	LoadHttp(app)

	log.Fatal(app.Listen())
}

func LoadHttp(app *core.App) {
	fb := app.Fiber()

	api := fb.Group("/api")

	uifs := ui.GetUIFS()

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

	// apidoc
	apidoc := fb.Group("/apidoc")
	h.Apidoc(app, apidoc)

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
	h.Feed(app, api)

	// site
	site := api.Group("/site")

	h.Site(app, site)

	// friend
	friend := api.Group("/friend")

	h.FriendList(app, friend)
	h.FriendAdd(app, friend)
	h.FriendUpdate(app, friend)
	h.FriendDelete(app, friend)

	// version
	h.Version(app, api)

	// dashboard
	dashboard(fb, uifs)

	// notFound
	h.NotFound(app, fb)

	articleCount, err := app.Query().Article.Count()

	if err != nil {
		log.Error("获取文章数量失败: ", err)
	}

	if articleCount == 0 {
		createFirstArticle(app)
	}
}

func loadPlugins(app *core.App) {
	plugins := app.Config().Plugins

	for _, pluginPath := range plugins {
		plugin.LoadPlugin(app, pluginPath)
	}
}

func dashboard(fb *fiber.App, uifs fs.FS) {
	fb.Get("/", func(c fiber.Ctx) error {
		path := "dist/index.html"

		file, err := uifs.Open(path)

		if err != nil {
			return common.RespServerError(c, err)
		}

		c.Type(".html")

		return c.SendStream(file)
	})

	fb.Get("/:path", func(c fiber.Ctx) error {
		path := "dist/" + c.Params("path")

		file, err := uifs.Open(path)

		if err != nil {
			if notFoundErr, ok := err.(*fs.PathError); ok && notFoundErr.Err == fs.ErrNotExist {
				return c.Next()
			}

			return common.RespServerError(c, err)
		}

		ext := filepath.Ext(path)
		c.Type(ext)

		return c.SendStream(file)
	})
}

func createFirstArticle(app *core.App) {
	err := app.Query().Article.Create(&model.Article{
		Slug:    "hello-world",
		Title:   "你好, 世界!",
		Desc:    "欢迎使用 reblog!",
		Content: "# 你好, 世界!\r\n\r\n欢迎使用 `reblog`，如果你能在文章列表看到这篇文章，那么说明reblog已经成功安装。\r\n",
	})

	if err != nil {
		log.Error("创建首篇文章失败: ", err)
	}
}
