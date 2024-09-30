package service

import "github.com/arklib/ark"

type (
	CreateIn  struct{}
	CreateOut struct{}
)

func (it *Service) UserCreate(c *ark.Ctx, in *CreateIn) (out *CreateOut, err error) {
	return
}
