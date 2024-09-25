package user

import (
	"github.com/arklib/ark"
	"github.com/arklib/ark/errx"
)

type (
	ApiLoginIn struct {
		Username string `json:"username" label:"用户名" vd:"required,min=4,max=32"`
		Password string `json:"password" label:"密码" vd:"required,min=6,max=32"`
	}
	ApiLoginOut struct {
		Token string `json:"token"`
	}
)

func (fn *Fn) ApiLogin(at *ark.At, in *ApiLoginIn) (out *ApiLoginOut, err error) {
	q := fn.Query.WithContext(at)
	userQuery := fn.Query.User

	user, err := q.User.Where(
		userQuery.Username.Eq(in.Username),
		userQuery.Password.Eq(in.Password),
	).First()
	errx.Assert(err, "auth failed")

	token, err := fn.Auth.NewToken(map[string]any{
		"type":   "user",
		"userId": user.Id,
	})
	errx.Assert(err, "create token failed")

	out = &ApiLoginOut{token}
	return
}
