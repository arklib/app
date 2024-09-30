package api

import (
	"github.com/spf13/cast"

	"github.com/arklib/ark"
	"github.com/arklib/ark/auth"
	"github.com/arklib/ark/errx"
)

type (
	TokenCreateIn struct {
		UserId uint `json:"userId"`
	}
	TokenCreateOut struct {
		UserId uint   `json:"userId"`
		Token  string `json:"token"`
	}
)

func (it *Api) TokenCreate(ctx *ark.Ctx, in *TokenCreateIn) (out *TokenCreateOut, err error) {
	payload := auth.Payload{
		"userId": in.UserId,
	}

	token, err := it.Auth.NewToken("user", payload)
	errx.Assert(err, "create token failed")
	it.Dump(token)

	// time.Sleep(time.Second * 2)
	claims, err := it.Auth.ParseToken(token)
	errx.Assert(err, "parse token failed")
	it.Dump(claims)

	userId := cast.ToUint(claims["userId"])
	out = &TokenCreateOut{userId, token}
	return
}
