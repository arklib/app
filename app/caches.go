package app

import (
	"demo/app/user/model"

	"github.com/arklib/ark/cache"
)

type Caches struct {
	KV   cache.Cache[string]
	User cache.Cache[model.User]
}

func (app *App) initCaches() {
	driver := cache.NewRedisDriver(app.Redis)

	caches := new(Caches)
	caches.KV = cache.Define[string](cache.Config{
		Driver: driver,
		Name:   "kv",
		TTL:    5 * 60,
	})

	caches.User = cache.Define[model.User](cache.Config{
		Driver: driver,
		Name:   "user",
		TTL:    10 * 60,
	})
	app.Caches = caches
}
