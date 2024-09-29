package api

import (
	"time"

	"github.com/arklib/ark"
)

type (
	PingIn  struct{}
	PingOut struct {
		Message string    `json:"message"`
		Time    time.Time `json:"time"`
	}
)

func (it *Api) Ping(at *ark.At, in *PingIn) (out *PingOut, err error) {
	out = &PingOut{"pong", time.Now()}
	return
}
