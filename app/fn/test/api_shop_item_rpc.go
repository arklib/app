package test

import (
	"github.com/arklib/ark"
	"github.com/arklib/ark/errx"

	"demo/app/hub/shop"
)

type (
	ApiShopItemRPCIn struct {
		Id int `json:"id"`
	}
	ApiShopItemRPCOut = shop.ApiShopItemGetOut
)

func (fn *Fn) ApiShopItemRPC(at *ark.At, in *ApiShopItemRPCIn) (out *ApiShopItemRPCOut, err error) {
	shopSvc := shop.New(fn.Server)
	out, err = shopSvc.ApiShopItemGet(at, &shop.ApiShopItemGetIn{
		Id: in.Id,
	})
	errx.Assert(err)
	return
}
