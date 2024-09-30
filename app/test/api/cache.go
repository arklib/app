package api

import (
	"time"

	"github.com/arklib/ark"
)

type (
	CacheIn  struct{}
	CacheOut struct {
		Time string `json:"time"`
	}
)

func (it *Api) Cache(ctx *ark.Ctx, in *CacheIn) (out *CacheOut, err error) {
	out = &CacheOut{}

	val, err := it.Caches.Any.Get(ctx, "time")
	if val != nil {
		out.Time = *val
		return
	}

	out.Time = time.Now().Format(time.DateTime)
	err = it.Caches.Any.Set(ctx, "time", &out.Time)
	return
}
