package top

import (
	"github.com/arklib/ark"
	"github.com/arklib/ark/errx"

	"app/src/dal/model"
)

type UserGetIn struct{}
type UserGetOut = model.User

// UserGet
func (it *Api) UserGet(at *ark.At, _ *UserGetIn) (out *UserGetOut, err error) {
	q := it.Query.WithContext(at)

	user, _ := it.Caches.User.Get(at, at.User.Id)
	if user != nil {
		user.Username = "(cache)" + user.Username
		return user, nil
	}

	user, err = q.User.Get(at.User.Id)
	errx.Assert(err, "search failed")

	out = user
	return
}
