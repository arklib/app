package test

import (
	"github.com/arklib/ark"
)

type (
	ApiValidateIn struct {
		Mobile string `json:"mobile" label:"手机" vd:"required"`
		Status string `json:"status" label:"状态" vd:"required,oneof='enable' 'disable'"`
	}
	ApiValidateOut struct {
		Mobile string `json:"mobile"`
		Status string `json:"status"`
	}
)

func (fn *Fn) ApiValidate(at *ark.At, in *ApiValidateIn) (out *ApiValidateOut, err error) {
	out = &ApiValidateOut{in.Mobile, in.Status}
	return
}
