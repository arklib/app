package api

import (
	"github.com/arklib/ark"
	"github.com/arklib/ark/errx"
)

type (
	LoginIn struct {
		Username string `json:"username" label:"用户名" vd:"required,min=4,max=32"`
		Password string `json:"password" label:"密码" vd:"required,min=6,max=32"`
	}
	LoginOut struct {
		Token string `json:"token"`
	}
)

func (it *Api) Login(c *ark.Ctx, in *LoginIn) (out *LoginOut, err error) {
	q := it.Query.WithContext(c)
	userQuery := it.Query.User

	user, err := q.User.Where(
		userQuery.Username.Eq(in.Username),
		userQuery.Password.Eq(in.Password),
	).First()
	errx.Assert(err, "auth failed")

	payload := map[string]any{
		"userId": user.Id,
	}
	token, err := it.Auth.NewToken("user", payload)
	errx.Assert(err, "create token failed")

	out = &LoginOut{token}
	return
}
