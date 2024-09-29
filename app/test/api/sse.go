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

func (it *Api) SSE(at *ark.At, in *SSEIn) (out *SSEOut, err error) {
	// Last-Emitter-Id
	lastEventID := sse.GetLastEventID(at.HttpCtx())
	fmt.Println(lastEventID)

	at.HttpCtx().SetStatusCode(http.StatusOK)
	stream := sse.NewStream(at.HttpCtx())
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
	at.HttpCtx().Abort()
	return
}
