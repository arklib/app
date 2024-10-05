package service

import (
	"context"
)

type (
	CreateIn  struct{}
	CreateOut struct{}
)

func (it *Service) UserCreate(ctx context.Context, in *CreateIn) (out *CreateOut, err error) {
	return
}
