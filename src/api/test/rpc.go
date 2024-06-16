package test

import (
	"github.com/arklib/ark"
	"github.com/arklib/ark/errx"

	"app/src/hub/shop"
)

type RPCIn struct{}
type RPCOut = shop.ItemGetOut

func (it *Api) RPC(at *ark.At, in *RPCIn) (out *RPCOut, err error) {
	shopSvc := shop.New(at)

	// client rpc
	out, err = shopSvc.ItemGet(&shop.ItemGetIn{
		Id: 1234,
	})
	errx.Assert(err)

	return
}
