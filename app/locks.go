package app

import (
	"github.com/arklib/ark/lock"
	"github.com/arklib/ark/lock/driver"
)

type Locks struct {
	User *lock.Lock
}

func (app *App) initLocks() {
	redisDriver := driver.NewRedisDriver(app.Redis)

	locks := new(Locks)
	locks.User = lock.Define(lock.Config{
		Name:   "user_create",
		TTL:    10,
		Driver: redisDriver,
	})
	app.Locks = locks
}
