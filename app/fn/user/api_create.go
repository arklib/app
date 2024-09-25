package user

import (
	"demo/app/model"
	"github.com/arklib/ark"
	"github.com/arklib/ark/errx"
)

type (
	ApiCreateIn struct {
		Nickname string `json:"nickname" label:"昵称" vd:"required,min=2,max=16"`
		Username string `json:"username" label:"用户名" vd:"required,min=4,max=32"`
		Password string `json:"password" label:"密码" vd:"required,min=6,max=32"`
	}
	ApiCreateOut = model.User
)

func (fn *Fn) ApiCreate(at *ark.At, in *ApiCreateIn) (out *ApiCreateOut, err error) {
	q := fn.Query.WithContext(at)
	
	ping, err := fn.test.ApiPing(at, nil)

	fn.Logger.Info(ping.Message)

	user := &model.User{
		Nickname: in.Nickname,
		Username: in.Username,
		Password: in.Password,
	}
	// event: user.create
	err = fn.Events.UserCreate.Emit(at, user)
	errx.Assert(err, "user create event failed")

	// db: user create
	err = q.User.Create(user)
	errx.Assert(err, "create failed")

	// cache: user
	err = fn.Caches.User.Set(at, user.Id, user)
	errx.Assert(err, "cache failed")

	out = user
	return
}
