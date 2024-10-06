package app

import (
	"demo/app/user/model"

	"github.com/arklib/ark/cache"
	"github.com/arklib/ark/cache/driver"
)

type Caches struct {
	KV   cache.Cache[string]
	User cache.Cache[model.User]
}

func (app *App) initCaches() {
	redisDriver := driver.NewRedisDriver(app.Redis)

	caches := new(Caches)
	caches.KV = cache.Define[string](cache.Config{
		Name:   "kv",
		TTL:    5 * 60,
		Driver: redisDriver,
	})

	caches.User = cache.Define[model.User](cache.Config{
		Name:   "user",
		TTL:    10 * 60,
		Driver: redisDriver,
	})
	app.Caches = caches
}
