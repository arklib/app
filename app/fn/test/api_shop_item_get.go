package test

import (
	"github.com/arklib/ark"
)

type (
	ApiShopItemGetIn struct {
		Id int `frugal:"1,default" json:"id" vd:"required"`
	}
	ApiShopItemGetOut struct {
		Id          int    `frugal:"1,default" json:"id"`
		Title       string `frugal:"2,default" json:"title"`
		Description string `frugal:"3,default" json:"description"`
	}
)

func (fn *Fn) ApiShopItemGet(at *ark.At, in *ApiShopItemGetIn) (out *ApiShopItemGetOut, err error) {
	out = &ApiShopItemGetOut{
		Id:          in.Id,
		Title:       "golang",
		Description: "go 1.22",
	}
	return
}
