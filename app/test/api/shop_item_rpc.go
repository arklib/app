package api

import (
	"demo/hub/shop"
	"github.com/arklib/ark"
	"github.com/arklib/ark/errx"
)

type (
	ShopItemRPCIn struct {
		Id int `json:"id"`
	}
	ShopItemRPCOut = shop.TestShopItemGetOut
)

func (it *Api) ShopItemRPC(ctx *ark.Ctx, in *ShopItemRPCIn) (out *ShopItemRPCOut, err error) {
	out, err = it.Shop.TestShopItemGet(ctx, &shop.TestShopItemGetIn{
		Id: in.Id,
	})
	errx.Assert(err)
	return
}
