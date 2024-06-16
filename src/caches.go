package src

import (
	"time"

	"github.com/arklib/ark/cache"

	"app/src/dal/model"
)

type Caches struct {
	Default cache.Cache[string]
	User    cache.Cache[model.User]
}

func initCaches(app *App) *Caches {
	driver := cache.NewRedis(app.Redis)
	return &Caches{
		Default: cache.New[string](
			driver,
			"common",
			1*time.Minute,
		),
		User: cache.New[model.User](
			driver,
			"user",
			10*time.Minute,
		),
	}
}
