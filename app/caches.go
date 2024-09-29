package app

import (
	"time"

	"demo/app/user/model"

	"github.com/arklib/ark/cache"
)

type Caches struct {
	Default cache.Cache[string]
	User    cache.Cache[model.User]
}

func (app *App) initCaches() {
	driver := cache.NewRedisDriver(app.UseRedis())
	app.Caches = &Caches{
		Default: cache.New[string](driver, "default", 1*time.Minute),
		User:    cache.New[model.User](driver, "user", 10*time.Minute),
	}
}
