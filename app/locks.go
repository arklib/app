package app

import (
	"time"

	"github.com/arklib/ark/lock"
)

type Locks struct {
	UserCreate *lock.Lock
}

func (app *App) initLocks() {
	driver := lock.NewRedisDriver(app.UseRedis())
	app.Locks = &Locks{
		UserCreate: lock.New(
			driver,
			"order:create",
			10*time.Second,
		),
	}
}
