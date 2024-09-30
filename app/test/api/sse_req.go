package api

import (
	"github.com/hertz-contrib/sse"

	"github.com/arklib/ark"
)

type (
	SSEReqIn  struct{}
	SSEReqOut struct{}
)

func (it *Api) SSEReq(c *ark.Ctx, in *SSEReqIn) (out *SSEReqOut, err error) {
	cli := sse.NewClient("http://127.0.0.1:8888/api/test/sse")

	evCh := make(chan *sse.Event)
	errCh := make(chan error)
	go func() {
		errCh <- cli.Subscribe(func(event *sse.Event) {
			if event.Data != nil {
				evCh <- event
				return
			}
		})
	}()
	for {
		select {
		case e := <-evCh:
			it.Logger.Infof("%s", e.Data)
		case err = <-errCh:
			if err != nil {
				it.Logger.Error(err)
			}
		}
	}
}
