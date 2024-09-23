package base

import (
	"gorm.io/gorm"

	"demo/app/query"
	"github.com/arklib/ark"
	"github.com/arklib/ark/auth"

	"github.com/redis/go-redis/v9"
)

type Base struct {
	*ark.Server
	fnMap map[string]any

	Auth  *auth.Auth
	DB    *gorm.DB
	Query *query.Query
	Redis redis.UniversalClient

	Locks  *Locks
	Caches *Caches
	Events *Events
}

func New(srv *ark.Server) *Base {
	base := &Base{
		Server: srv,
		fnMap:  make(map[string]any),
	}
	base.initEvents()
	return base
}

func (base *Base) Init() *Base {
	base.initAuth()
	base.initDB()
	base.initRedis()
	base.initLocks()
	base.initCaches()
	return base
}

func (base *Base) Use(handlers ...func(*Base)) *Base {
	for _, handler := range handlers {
		handler(base)
	}
	return base
}

func (base *Base) Fn(name string) any {
	return base.fnMap[name]
}

func (base *Base) AddFn(name string, fn any) {
	base.fnMap[name] = fn
}
