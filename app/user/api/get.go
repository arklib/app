package api

import (
	"demo/app/user/model"

	"github.com/arklib/ark"
	"github.com/arklib/ark/errx"
)

type (
	GetIn struct {
		Id uint `auth:"userId"`
	}
	GetOut = model.User
)

func (it *Api) Get(at *ark.At, in *GetIn) (out *GetOut, err error) {
	q := it.Query.WithContext(at)

	user, _ := it.Caches.User.Get(at, in.Id)
	if user != nil {
		user.Username = "(cache)" + user.Username
		return user, nil
	}

	user, err = q.User.Get(in.Id)
	errx.Assert(err, "search failed")

	out = user
	return
}
