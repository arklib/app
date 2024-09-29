package api

import (
	"github.com/arklib/ark"
)

type (
	ShopItemGetIn struct {
		Id int `frugal:"1,default" json:"id" vd:"required"`
	}
	ShopItemGetOut struct {
		Id          int    `frugal:"1,default" json:"id"`
		Title       string `frugal:"2,default" json:"title"`
		Description string `frugal:"3,default" json:"description"`
	}
)

func (it *Api) ShopItemGet(at *ark.At, in *ShopItemGetIn) (out *ShopItemGetOut, err error) {
	out = &ShopItemGetOut{
		Id:          in.Id,
		Title:       "golang",
		Description: "go 1.22",
	}
	return
}
