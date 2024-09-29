package service

import "github.com/arklib/ark"

type (
	CreateIn  struct{}
	CreateOut struct{}
)

func (it *Service) UserCreate(at *ark.At, in *CreateIn) (out *CreateOut, err error) {
	return
}
