package test

import (
	"time"

	"github.com/arklib/ark"
)

type LockIn struct{}
type LockOut struct {
	Message string `json:"message"`
}

func (it *Api) Lock(at *ark.At, in *LockIn) (out *LockOut, err error) {
	lock, err := it.Locks.UserOrderCreate.Apply(at, "id")
	if err != nil {
		return
	}
	defer func() { err = lock.Release() }()

	time.Sleep(10 * time.Second)
	out = &LockOut{"ok"}
	return
}
