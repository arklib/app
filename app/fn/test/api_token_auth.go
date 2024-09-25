package test

import (
	"github.com/arklib/ark"
)

type (
	ApiTokenAuthIn  struct{}
	ApiTokenAuthOut struct {
		UserId uint `json:"userId"`
	}
)

func (fn *Fn) ApiTokenAuth(at *ark.At, in *ApiTokenAuthIn) (out *ApiTokenAuthOut, err error) {
	out = &ApiTokenAuthOut{
		1,
	}
	return
}
