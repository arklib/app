package app

import (
	"github.com/arklib/ark/lock"
)

type Locks struct {
	User *lock.Lock
}

func (app *App) initLocks() {
	driver := lock.NewRedisDriver(app.Redis)

	locks := new(Locks)
	locks.User = lock.Define(lock.Config{
		Driver: driver,
		Name:   "user_create",
		TTL:    10,
	})
	app.Locks = locks
}
