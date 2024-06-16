package shop

import (
	"github.com/arklib/ark"
)

type ItemGetIn struct {
	Id int `frugal:"1,default" json:"id" vd:"required"`
}

type ItemGetOut struct {
	Id          int    `frugal:"1,default" json:"id"`
	Title       string `frugal:"2,default" json:"title"`
	Description string `frugal:"3,default" json:"description"`
}

func (it *Api) ItemGet(at *ark.At, in *ItemGetIn) (out *ItemGetOut, err error) {
	out = &ItemGetOut{
		Id:          in.Id,
		Title:       "golang",
		Description: "go 1.22",
	}
	return
}
