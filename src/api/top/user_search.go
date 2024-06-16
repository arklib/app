package top

import (
	"github.com/arklib/ark"
	"github.com/arklib/ark/errx"

	"app/src/dal/model"
)

type UserSearchIn struct {
	Page int `json:"page" label:"页码" vd:"required,min=1"`
	Size int `json:"size" label:"页面大小" vd:"required,min=10,max=100"`
}
type UserSearchOut struct {
	List []*model.User `json:"list"`
}

// UserGet
func (it *Api) UserSearch(at *ark.At, in *UserSearchIn) (out *UserSearchOut, err error) {
	q := it.Query.WithContext(at)

	offset := in.Size * (in.Page - 1)
	users, err := q.User.Offset(offset).Limit(in.Size).Find()
	errx.Assert(err, "search failed")

	out = &UserSearchOut{List: users}
	return
}
