package src

import (
	"github.com/arklib/ark"

	"github.com/arklib/ark/auth"

	"gorm.io/gorm"

	"github.com/redis/go-redis/v9"

	"app/src/dal/query"
	"app/src/lib"
)

type App struct {
	*ark.Server
	Events *Events
	Auth   *auth.Auth

	DB    *gorm.DB
	Query *query.Query

	Redis  redis.UniversalClient
	Locks  *Locks
	Caches *Caches

	// lib
	Email *lib.Email
	SMS   *lib.SMS
}

func NewApp(srv *ark.Server) *App {
	app := new(App)
	app.Server = srv

	app.DB = initDB(app)
	app.Query = query.Use(app.DB)
	app.Redis = initRedis(app)
	app.Auth = initAuth(app)

	// lib
	app.Email = lib.NewEmail(srv)
	app.SMS = lib.NewSMS(srv)

	app.Locks = initLocks(app)
	app.Caches = initCaches(app)
	app.Events = initEvents(app)
	return app
}

func (app *App) Use(handlers ...func(*App)) {
	for _, handler := range handlers {
		handler(app)
	}
}
