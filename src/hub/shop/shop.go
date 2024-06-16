// Code generated by ark. DO NOT EDIT.

package shop

import "github.com/arklib/ark"

type ItemGetIn struct {
	Id int `frugal:"1,default" json:"id" vd:"required"`
}

type ItemGetOut struct {
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

func (s *Service) ItemGet(in *ItemGetIn) (out *ItemGetOut, err error) {
	out = new(ItemGetOut)
	err = s.at.FetchSvc("shop/item/get", in, out)
	return
}
