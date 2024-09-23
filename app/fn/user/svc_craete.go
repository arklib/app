package user

import "context"

type (
	CreateIn  struct{}
	CreateOut struct{}
)

func (fn *Fn) Create(ctx context.Context, in *CreateIn) (out *CreateOut, err error) {
	return
}
