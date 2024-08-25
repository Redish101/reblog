// reblog vercel 兼容层

package handler

import (
	"net/http"

	"github.com/redish101/reblog/internal/config"
	"github.com/redish101/reblog/internal/core"
	"github.com/redish101/reblog/server"

	"github.com/gofiber/fiber/v3/middleware/adaptor"
)

var srv http.Handler

func init() {
	config := config.NewFromFile()
	app := core.NewApp(config)
	app.Bootstrap()

	server.LoadHttp(app)

	srv = adaptor.FiberApp(app.Fiber())
}

func Handler(w http.ResponseWriter, r *http.Request) {
	srv.ServeHTTP(w, r)
}
