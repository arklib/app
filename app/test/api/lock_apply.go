package api

import (
	"time"

	"github.com/arklib/ark"
)

type (
	LockApplyIn  struct{}
	LockApplyOut struct {
		Message string `json:"message"`
	}
)

func (it *Api) LockApply(at *ark.At, in *LockApplyIn) (out *LockApplyOut, err error) {
	lock, err := it.Locks.UserCreate.Lock(at, "id")
	if err != nil {
		return
	}
	defer func() { err = lock.Free() }()

	time.Sleep(10 * time.Second)
	out = &LockApplyOut{"ok"}
	return
}
