// Code generated by ark. DO NOT EDIT.

package shop

import "github.com/arklib/ark"

type TestShopItemGetIn struct {
	Id int `frugal:"1,default" json:"id" vd:"required"`
}

type TestShopItemGetOut struct {
	Id          int    `frugal:"1,default" json:"id"`
	Title       string `frugal:"2,default" json:"title"`
	Description string `frugal:"3,default" json:"description"`
}

type Service struct {
	srv *ark.Server
}

func New(srv *ark.Server) *Service {
	return &Service{srv}
}

func (s *Service) TestShopItemGet(ctx *ark.Ctx, in *TestShopItemGetIn) (out *TestShopItemGetOut, err error) {
	out = new(TestShopItemGetOut)
	err = s.srv.RPC(ctx, "shop/test/get", in, out)
	return
}