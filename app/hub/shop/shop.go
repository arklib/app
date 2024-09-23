// Code generated by ark. DO NOT EDIT.

package shop

import "github.com/arklib/ark"

type ApiShopItemGetIn struct {
	Id int `frugal:"1,default" json:"id" vd:"required"`
}

type ApiShopItemGetOut struct {
	Id          int    `frugal:"1,default" json:"id"`
	Title       string `frugal:"2,default" json:"title"`
	Description string `frugal:"3,default" json:"description"`
}

type Service struct {
	at *ark.At
}

func New(at *ark.At) *Service {
	return &Service{at}
}

func (s *Service) ApiShopItemGet(in *ApiShopItemGetIn) (out *ApiShopItemGetOut, err error) {
	out = new(ApiShopItemGetOut)
	err = s.at.FetchSvc("shop/item/get", in, out)
	return
}