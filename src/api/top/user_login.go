package top

import (
	"github.com/arklib/ark"
	"github.com/arklib/ark/auth"
	"github.com/arklib/ark/errx"
)

type UserLoginIn struct {
	Username string `json:"username" label:"用户名" vd:"required,min=4,max=32"`
	Password string `json:"password" label:"密码" vd:"required,min=6,max=32"`
}
type UserLoginOut struct {
	Token string `json:"token"`
}

// UserLogin
func (it *Api) UserLogin(at *ark.At, in *UserLoginIn) (out *UserLoginOut, err error) {
	q := it.Query.WithContext(at)
	u := it.Query.User

	user, err := q.User.Where(
		u.Username.Eq(in.Username),
		u.Password.Eq(in.Password),
	).First()
	errx.Assert(err, "auth failed")

	token, err := it.Auth.NewToken(&auth.User{
		Id:   user.Id,
		Role: "user",
	})
	errx.Assert(err, "create token failed")

	out = &UserLoginOut{token}
	return
}
