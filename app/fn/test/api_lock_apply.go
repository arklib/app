package test

import (
	"time"

	"github.com/arklib/ark"
)

type (
	ApiLockApplyIn  struct{}
	ApiLockApplyOut struct {
		Message string `json:"message"`
	}
)

func (fn *Fn) ApiLockApply(at *ark.At, in *ApiLockApplyIn) (out *ApiLockApplyOut, err error) {
	lock, err := fn.Locks.UserCreate.Apply(at, "id")
	if err != nil {
		return
	}
	defer func() { err = lock.Release() }()

	time.Sleep(10 * time.Second)
	out = &ApiLockApplyOut{"ok"}
	return
}
