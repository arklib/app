package top

import (
	"github.com/arklib/ark"
	"github.com/arklib/ark/errx"

	"app/src/dal/model"
)

type UserCreateIn struct {
	Nickname string `json:"nickname" label:"昵称" vd:"required,min=2,max=16"`
	Username string `json:"username" label:"用户名" vd:"required,min=4,max=32"`
	Password string `json:"password" label:"密码" vd:"required,min=6,max=32"`
}
type UserCreateOut = model.User

// UserCreate
func (it *Api) UserCreate(at *ark.At, in *UserCreateIn) (out *UserCreateOut, err error) {
	q := it.Query.WithContext(at)

	user := &model.User{
		Nickname: in.Nickname,
		Username: in.Username,
		Password: in.Password,
	}
	// event: user.create
	err = it.Events.UserCreate(at, user)
	errx.Assert(err, "user create event failed")

	// db: user create
	err = q.User.Create(user)
	errx.Assert(err, "create failed")

	// cache: user
	err = it.Caches.User.Set(at, user.Id, user)
	errx.Assert(err, "cache failed")

	out = user
	return
}
