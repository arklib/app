package base

import (
	"gorm.io/gorm"

	"github.com/redis/go-redis/v9"

	"demo/app/model/query"
	"github.com/arklib/ark"
	"github.com/arklib/ark/auth"
)

type Base struct {
	*ark.Server
	fnMap map[string]any

	Auth  *auth.Auth
	DB    *gorm.DB
	Query *query.Query
	Redis redis.UniversalClient

	Locks  *Locks
	Tasks  *Tasks
	Caches *Caches
	Events *Events
}

func New(srv *ark.Server) *Base {
	base := &Base{
		Server: srv,
		fnMap:  make(map[string]any),
	}
	base.defineEvents()
	return base
}

func (base *Base) Init() *Base {
	base.GetAuth()
	base.GetDB()
	base.GetRedis()
	base.GetLocks()
	base.GetTasks()
	base.GetCaches()
	return base
}

func (base *Base) GetFn(name string) any {
	return base.fnMap[name]
}

func (base *Base) AddFn(name string, fn any) *Base {
	base.fnMap[name] = fn
	return base
}

func (base *Base) Use(handlers ...func(*Base)) *Base {
	for _, handler := range handlers {
		handler(base)
	}
	return base
}
