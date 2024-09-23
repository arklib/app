package user

import "github.com/arklib/ark"

type (
	CreateIn  struct{}
	CreateOut struct{}
)

func (fn *Fn) Create(at *ark.At, in *CreateIn) (out *CreateOut, err error) {
	return
}
