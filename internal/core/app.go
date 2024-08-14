package core

import (
	"fmt"

	"github.com/ChuqiCloud/acmeidc/internal/config"
	"github.com/ChuqiCloud/acmeidc/internal/db"
	"github.com/ChuqiCloud/acmeidc/internal/log"
	"github.com/ChuqiCloud/acmeidc/internal/query"
	"github.com/ChuqiCloud/acmeidc/internal/version"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"github.com/sirupsen/logrus"
)

type App struct {
	config    *config.Config
	fiber     *fiber.App
	query     *query.Query
	validator *validator.Validate
	dev       bool

	service *map[string]Service
}

// 注入服务到App实例
func (a *App) Inject(name string, service Service) {
	(*a.service)[name] = service
}

// 注入服务到App实例, 并生成服务名称
func AppInject[T Service](app *App, service T) {
	log.Debugf("[SERVICE] 注入服务 %s", getServiceName[T]())
	app.Inject(getServiceName[T](), service)
}

func (app *App) Service(name string) (Service, error) {
	if app.service == nil {
		return nil, fmt.Errorf("服务未初始化")
	}

	if _, isExits := (*app.service)[name]; !isExits {
		return nil, fmt.Errorf("服务 %s 不存在", name)
	}

	return (*app.service)[name], nil
}

func AppService[T Service](app *App) (T, error) {
	service, err := app.Service(getServiceName[T]())

	if err != nil {
		var zero T
		return zero, err
	}

	return service.(T), nil
}

func (app *App) initFiber() {
	app.fiber = GetFiber()
}

func (app *App) initQuery() {
	query.Use(db.DB(&app.config.DB))
	query.SetDefault(db.DB(&app.config.DB))

	app.query = query.Q
}

func (app *App) initValidator() {
	app.validator = validator.New()
}

func (app *App) initDefaultServices() {
	AppInject(app, NewAuthService(app))
}

func (app *App) Bootstrap() {
	for i := range *app.service {
		if err := (*app.service)[i].Start(); err != nil {
			log.Errorf("服务 %s 启动失败: %s", i, err)
		}
	}
}

func (app *App) Config() *config.Config {
	return app.config
}

func (app *App) Fiber() *fiber.App {
	return app.fiber
}

func (app *App) Query() *query.Query {
	return app.query
}

func (app *App) Validator() *validator.Validate {
	return app.validator
}

func (app *App) Dev() bool {
	return app.dev
}

func (app *App) Listen() error {
	serverConfig := app.config.Server

	host := serverConfig.Host
	port := serverConfig.Port

	listenConfig := fiber.ListenConfig{
		DisableStartupMessage: true,
	}

	listenUrl := fmt.Sprintf("%s:%d", host, port)

	log.Infof("[HTTP] 在 http://%s 启动服务", listenUrl)

	return app.fiber.Listen(listenUrl, listenConfig)
}

func NewApp(config *config.Config) *App {
	app := &App{
		config:  config,
		service: &map[string]Service{},
	}

	if version.Version == "dev" {
		app.dev = true
		log.Logger().SetLevel(logrus.DebugLevel)
		log.Debug("以开发模式启动")
	} else {
		app.dev = false
	}

	app.initFiber()
	app.initQuery()
	app.initValidator()

	app.initDefaultServices()

	return app
}
