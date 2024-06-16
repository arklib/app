package test

import (
	"time"

	"github.com/arklib/ark"
)

type CacheIn struct{}
type CacheOut struct {
	Value string `json:"message"`
}

func (it *Api) Cache(at *ark.At, in *CacheIn) (out *CacheOut, err error) {
	out = &CacheOut{}

	val, err := it.Caches.Default.Get(at, "time")
	if val != nil {
		out.Value = *val
		return
	}

	out.Value = time.Now().Format(time.DateTime)
	err = it.Caches.Default.Set(at, "time", &out.Value)
	return
}
