package test

import (
	"github.com/spf13/cast"

	"github.com/arklib/ark"
	"github.com/arklib/ark/auth"
	"github.com/arklib/ark/errx"
)

type (
	ApiTokenCreateIn struct {
		UserId uint `json:"userId"`
	}
	ApiTokenCreateOut struct {
		UserId uint   `json:"userId"`
		Token  string `json:"token"`
	}
)

func (fn *Fn) ApiTokenCreate(at *ark.At, in *ApiTokenCreateIn) (out *ApiTokenCreateOut, err error) {
	payload := auth.Payload{
		"userId": in.UserId,
	}

	token, err := fn.Auth.NewToken("user", payload)
	errx.Assert(err, "create token failed")
	fn.Dump(token)

	// time.Sleep(time.Second * 2)
	claims, err := fn.Auth.ParseToken(token)
	errx.Assert(err, "parse token failed")
	fn.Dump(claims)

	userId := cast.ToUint(claims["userId"])
	out = &ApiTokenCreateOut{userId, token}
	return
}
