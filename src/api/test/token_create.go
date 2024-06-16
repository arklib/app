package test

import (
	"github.com/arklib/ark"
	"github.com/arklib/ark/auth"
	"github.com/arklib/ark/errx"
)

type TokenCreateIn struct{}
type TokenCreateOut struct {
	UserId uint   `json:"userId"`
	Token  string `json:"token"`
}

func (it *Api) TokenCreate(at *ark.At, in *TokenCreateIn) (out *TokenCreateOut, err error) {
	user := &auth.User{
		Id:   123456,
		Role: "user",
	}
	token, err := it.Auth.NewToken(user)
	errx.Assert(err, "create token failed")

	// time.Sleep(time.Second)
	authUser, err := it.Auth.ParseToken(token)
	errx.Assert(err, "parse token failed")

	it.Dump(authUser)

	out = &TokenCreateOut{authUser.Id, token}
	return
}
