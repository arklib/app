package base

import (
	"time"

	"github.com/arklib/ark/lock"
)

type Locks struct {
	UserCreate *lock.Lock
}

func (base *Base) GetLocks() *Locks {
	if base.Locks != nil {
		return base.Locks
	}

	driver := lock.NewRedisDriver(base.GetRedis())
	base.Locks = &Locks{
		UserCreate: lock.New(driver, "order:create", 10*time.Second),
	}
	return base.Locks
}
