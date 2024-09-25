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

func (base *Base) GetCaches() *Caches {
	if base.Caches != nil {
		return base.Caches
	}

	driver := cache.NewRedisDriver(base.GetRedis())
	base.Caches = &Caches{
		Default: cache.New[string](driver, "default", 1*time.Minute),
		User:    cache.New[model.User](driver, "user", 10*time.Minute),
	}
	return base.Caches
}
