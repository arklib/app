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
	models map[string]any

	DB     *gorm.DB
	Query  *query.Query
	Redis  redis.UniversalClient
	Auth   *auth.Auth
	Locks  *Locks
	Caches *Caches
	Queues *Queues
	Hooks  *Hooks

	Shop *shop.Service
}

func New(srv *ark.Server) *App {
	app := &App{
		Server: srv,
		models: make(map[string]any),
	}
	return app.init()
}

func (app *App) init() *App {
	app.initDB()
	app.initRedis()
	app.initAuth()
	app.initLocks()
	app.initCaches()
	app.initQueues()
	app.initHooks()
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
