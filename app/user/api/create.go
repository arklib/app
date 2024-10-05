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

func (it *Api) Create(ctx *ark.Ctx, in *CreateIn) (out *CreateOut, err error) {
	q := it.Query.WithContext(ctx)

	ping, err := it.testApi.Ping(ctx, nil)

	it.Logger.Info(ping.Message)

	user := &model.User{
		Nickname: in.Nickname,
		Username: in.Username,
		Password: in.Password,
	}
	err = q.User.Create(user)
	errx.Assert(err, "create failed")

	// dispatch user create
	err = it.Hooks.UserCreateAfter.Emit(ctx, user)
	errx.Assert(err, "user create hook failed")

	// cache: user
	err = it.Caches.User.Set(ctx, user.ID, user)
	errx.Assert(err, "cache failed")

	out = user
	return
}
