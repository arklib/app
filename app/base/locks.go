package base

import (
	"time"

	"github.com/arklib/ark/lock"
)

type Locks struct {
	UserCreate *lock.Lock
}

func (base *Base) initLocks() {
	driver := lock.NewRedisDriver(base.Redis)
	base.Locks = &Locks{
		UserCreate: lock.New(driver, "order:create", 10*time.Second),
	}
}
