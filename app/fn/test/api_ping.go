package test

import (
	"time"

	"github.com/arklib/ark"
)

type (
	ApiPingIn  struct{}
	ApiPingOut struct {
		Message string    `json:"message"`
		Time    time.Time `json:"time"`
	}
)

func (fn *Fn) ApiPing(at *ark.At, in *ApiPingIn) (out *ApiPingOut, err error) {
	out = &ApiPingOut{"pong", time.Now()}
	return
}
