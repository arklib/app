package user

import (
	"demo/app/model"

	"github.com/arklib/ark"
	"github.com/arklib/ark/errx"
)

type (
	ApiGetIn struct {
		Id uint `auth:"userId"`
	}
	ApiGetOut = model.User
)

func (fn *Fn) ApiGet(at *ark.At, in *ApiGetIn) (out *ApiGetOut, err error) {
	q := fn.Query.WithContext(at)

	user, _ := fn.Caches.User.Get(at, in.Id)
	if user != nil {
		user.Username = "(cache)" + user.Username
		return user, nil
	}

	user, err = q.User.Get(in.Id)
	errx.Assert(err, "search failed")

	out = user
	return
}
