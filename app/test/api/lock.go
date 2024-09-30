package api

import (
	"time"

	"github.com/arklib/ark"
)

type (
	LockIn struct {
		UserId uint `json:"id"`
	}
	LockOut struct {
		Message string `json:"message"`
	}
)

func (it *Api) Lock(ctx *ark.Ctx, in *LockIn) (out *LockOut, err error) {
	lock, err := it.Locks.User.Lock(ctx, in.UserId)
	if err != nil {
		return
	}
	defer func() { err = lock.Free() }()

	time.Sleep(10 * time.Second)
	out = &LockOut{"ok"}
	return
}
