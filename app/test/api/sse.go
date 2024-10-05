package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/hertz-contrib/sse"

	"github.com/arklib/ark"
)

type (
	SSEIn  struct{}
	SSEOut struct{}
)

func (it *Api) SSE(ctx *ark.Ctx, in *SSEIn) (out *SSEOut, err error) {
	req := ctx.Http()

	// Last-Emitter-Id
	lastEventID := sse.GetLastEventID(req)
	fmt.Println(lastEventID)

	req.SetStatusCode(http.StatusOK)
	stream := sse.NewStream(req)
	i := 0
	for range time.NewTicker(1 * time.Second).C {
		if i == 3 {
			break
		}

		data, _ := json.Marshal(map[string]any{
			"id":   i + 1,
			"date": time.Now().Format(time.DateTime),
		})
		event := &sse.Event{
			Event: "timestamp",
			Data:  data,
		}

		err = stream.Publish(event)
		if err != nil {
			break
		}
		i++
	}
	req.Abort()
	return
}
