package test

import (
	"github.com/arklib/ark"
)

type ValidateIn struct {
	Mobile string `json:"mobile" label:"手机" vd:"required"`
	Status string `json:"status" label:"状态" vd:"required,oneof='enable' 'disable'"`
}
type ValidateOut struct {
	Mobile string `json:"mobile"`
	Status string `json:"status"`
}

func (it *Api) Validate(at *ark.At, in *ValidateIn) (out *ValidateOut, err error) {
	out = &ValidateOut{in.Mobile, in.Status}
	return
}
