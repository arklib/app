package app

import (
	"demo/app/user/model"

	"github.com/arklib/ark/cache"
)

type Caches struct {
	Any  cache.Cache[string]
	User cache.Cache[model.User]
}

func (app *App) initCaches() {
	driver := cache.NewRedisDriver(app.GetRedis())

	app.Caches = &Caches{
		Any: cache.Define[string](cache.Config{
			Driver: driver,
			Scene:  "any",
			TTL:    5 * 60,
		}),
		User: cache.Define[model.User](cache.Config{
			Driver: driver,
			Scene:  "user",
			TTL:    10 * 60,
		}),
	}
}
