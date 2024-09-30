package api

import (
	"github.com/arklib/ark"
)

type (
	TokenAuthIn struct {
		UserId uint `auth:"userId"`
	}
	TokenAuthOut struct {
		UserId uint `json:"userId"`
	}
)

func (it *Api) TokenAuth(ctx *ark.Ctx, in *TokenAuthIn) (out *TokenAuthOut, err error) {
	out = &TokenAuthOut{
		in.UserId,
	}
	return
}
