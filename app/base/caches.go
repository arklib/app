package base

import (
	"time"

	"demo/app/model"
	"github.com/arklib/ark/cache"
)

type Caches struct {
	Default cache.Cache[string]
	User    cache.Cache[model.User]
}

func (base *Base) initCaches() {
	driver := cache.NewRedisDriver(base.Redis)
	base.Caches = &Caches{
		Default: cache.New[string](driver, "default", 1*time.Minute),
		User:    cache.New[model.User](driver, "user", 10*time.Minute),
	}
}
