package test

import (
	"github.com/arklib/ark"
	"github.com/arklib/ark/auth"
	"github.com/arklib/ark/errx"
)

type (
	ApiTokenCreateIn  struct{}
	ApiTokenCreateOut struct {
		UserId uint   `json:"userId"`
		Token  string `json:"token"`
	}
)

func (fn *Fn) ApiTokenCreate(at *ark.At, in *ApiTokenCreateIn) (out *ApiTokenCreateOut, err error) {
	user := &auth.User{
		Id:   123456,
		Role: "user",
	}
	token, err := fn.Auth.NewToken(user)
	errx.Assert(err, "create token failed")

	// time.Sleep(time.Second)
	authUser, err := fn.Auth.ParseToken(token)
	errx.Assert(err, "parse token failed")

	fn.Dump(authUser)

	out = &ApiTokenCreateOut{authUser.Id, token}
	return
}
