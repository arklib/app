package user

import (
	"github.com/arklib/ark"
	"github.com/arklib/ark/errx"

	"demo/app/model"
)

type (
	ApiGetIn  struct{}
	ApiGetOut = model.User
)

func (fn *Fn) ApiGet(at *ark.At, _ *ApiGetIn) (out *ApiGetOut, err error) {
	q := fn.Query.WithContext(at)

	user, _ := fn.Caches.User.Get(at, at.User.Id)
	if user != nil {
		user.Username = "(cache)" + user.Username
		return user, nil
	}

	user, err = q.User.Get(at.User.Id)
	errx.Assert(err, "search failed")

	out = user
	return
}
