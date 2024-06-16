package src

import (
	"time"

	"github.com/arklib/ark/lock"
)

type Locks struct {
	UserOrderCreate *lock.Lock
}

func initLocks(app *App) *Locks {
	driver := lock.NewRedis(app.Redis)
	return &Locks{
		UserOrderCreate: lock.New(
			driver,
			"order.create",
			10*time.Second,
		),
	}
}
