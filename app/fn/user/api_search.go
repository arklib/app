package user

import (
	"github.com/arklib/ark"
	"github.com/arklib/ark/errx"

	"demo/app/model"
)

type (
	ApiSearchIn struct {
		Page int `json:"page" label:"页码" vd:"required,min=1"`
		Size int `json:"size" label:"页面大小" vd:"required,min=10,max=100"`
	}
	SearchOut struct {
		List []*model.User `json:"list"`
	}
)

func (fn *Fn) ApiSearch(at *ark.At, in *ApiSearchIn) (out *SearchOut, err error) {
	q := fn.Query.WithContext(at)

	offset := in.Size * (in.Page - 1)
	users, err := q.User.Offset(offset).Limit(in.Size).Find()
	errx.Assert(err, "search failed")

	out = &SearchOut{List: users}
	return
}
