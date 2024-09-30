package app

import (
	"github.com/arklib/ark/lock"
)

type Locks struct {
	User *lock.Lock
}

func (app *App) initLocks() {
	driver := lock.NewRedisDriver(app.Redis)

	app.Locks = &Locks{
		User: lock.Define(lock.Config{
			Driver: driver,
			Scene:  "user:create",
			TTL:    10,
		}),
	}
}
