package test

import (
	"github.com/arklib/ark"
)

type TokenAuthIn struct{}
type TokenAuthOut struct {
	UserId uint `json:"userId"`
}

func (it *Api) TokenAuth(at *ark.At, in *TokenAuthIn) (out *TokenAuthOut, err error) {
	out = &TokenAuthOut{
		at.User.Id,
	}
	return
}
