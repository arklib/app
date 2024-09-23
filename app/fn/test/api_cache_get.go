package test

import (
	"time"

	"github.com/arklib/ark"
)

type (
	ApiCacheGetIn  struct{}
	ApiCacheGetOut struct {
		Value string `json:"message"`
	}
)

func (fn *Fn) ApiCacheGet(at *ark.At, in *ApiCacheGetIn) (out *ApiCacheGetOut, err error) {
	out = &ApiCacheGetOut{}

	val, err := fn.Caches.Default.Get(at, "time")
	if val != nil {
		out.Value = *val
		return
	}

	out.Value = time.Now().Format(time.DateTime)
	err = fn.Caches.Default.Set(at, "time", &out.Value)
	return
}
