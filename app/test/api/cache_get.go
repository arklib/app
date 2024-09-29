package api

import (
	"time"

	"github.com/arklib/ark"
)

type (
	CacheGetIn  struct{}
	CacheGetOut struct {
		Value string `json:"message"`
	}
)

func (it *Api) CacheGet(at *ark.At, in *CacheGetIn) (out *CacheGetOut, err error) {
	out = &CacheGetOut{}

	val, err := it.Caches.Default.Get(at, "time")
	if val != nil {
		out.Value = *val
		return
	}

	out.Value = time.Now().Format(time.DateTime)
	err = it.Caches.Default.Set(at, "time", &out.Value)
	return
}
