package api

import (
	"demo/app/user/model"

	"github.com/arklib/ark"
	"github.com/arklib/ark/errx"
)

type (
	CreateIn struct {
		Nickname string `json:"nickname" label:"昵称" vd:"required,min=2,max=16"`
		Username string `json:"username" label:"用户名" vd:"required,min=4,max=32"`
		Password string `json:"password" label:"密码" vd:"required,min=6,max=32"`
	}
	CreateOut = model.User
)

func (it *Api) Create(c *ark.Ctx, in *CreateIn) (out *CreateOut, err error) {
	q := it.Query.WithContext(c)

	ping, err := it.testApi.Ping(c, nil)

	it.Logger.Info(ping.Message)

	user := &model.User{
		Nickname: in.Nickname,
		Username: in.Username,
		Password: in.Password,
	}
	// event: user.create
	err = it.Events.UserCreate.Dispatch(c, user)
	errx.Assert(err, "user create event failed")

	// db: user create
	err = q.User.Create(user)
	errx.Assert(err, "create failed")

	// cache: user
	err = it.Caches.User.Set(c, user.Id, user)
	errx.Assert(err, "cache failed")

	out = user
	return
}
