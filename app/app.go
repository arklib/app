package app

import (
	"gorm.io/gorm"

	"github.com/redis/go-redis/v9"

	"demo/etc/query"
	"demo/hub/shop"
	"github.com/arklib/ark"
	"github.com/arklib/ark/auth"
)

type App struct {
	*ark.Server
	apis     map[string]any
	models   map[string]any
	services map[string]any

	DB     *gorm.DB
	Query  *query.Query
	Redis  redis.UniversalClient
	Auth   *auth.Auth
	Events *Events
	Jobs   *Jobs
	Locks  *Locks
	Caches *Caches

	Shop *shop.Service
}

func New(srv *ark.Server) *App {
	app := &App{
		Server:   srv,
		apis:     make(map[string]any),
		models:   make(map[string]any),
		services: make(map[string]any),
	}
	return app.init()
}

func (app *App) init() *App {
	app.initDB()
	app.initRedis()
	app.initAuth()
	app.initEvents()
	app.initJobs()
	app.initLocks()
	app.initCaches()

	app.Shop = shop.New(app.Server)
	return app
}

func (app *App) Run() *App {
	app.init().Server.Run()
	return app
}

func (app *App) Use(handlers ...func(*App)) *App {
	for _, handler := range handlers {
		handler(app)
	}
	return app
}

func (app *App) AddApi(name string, instance any) *App {
	app.apis[name] = instance
	return app
}

func (app *App) GetApi(name string) any {
	return app.apis[name]
}

func (app *App) AddService(name string, instance any) *App {
	app.services[name] = instance
	return app
}

func (app *App) GetService(name string) any {
	return app.services[name]
}

func (app *App) AddModel(name string, instance any) *App {
	app.models[name] = instance
	return app
}

func (app *App) GetModel(name string) any {
	return app.models[name]
}

func (app *App) GetModels() []any {
	var models []any
	for _, model := range app.models {
		models = append(models, model)
	}
	return models
}
