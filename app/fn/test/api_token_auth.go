package test

import (
	"github.com/arklib/ark"
)

type (
	ApiTokenAuthIn struct {
		UserId uint `auth:"userId"`
	}
	ApiTokenAuthOut struct {
		UserId uint `json:"userId"`
	}
)

func (fn *Fn) ApiTokenAuth(at *ark.At, in *ApiTokenAuthIn) (out *ApiTokenAuthOut, err error) {
	out = &ApiTokenAuthOut{
		in.UserId,
	}
	return
}
