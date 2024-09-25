package test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/hertz-contrib/sse"

	"github.com/arklib/ark"
)

type (
	ApiSSEIn  struct{}
	ApiSSEOut struct{}
)

func (fn *Fn) ApiSSE(at *ark.At, in *ApiSSEIn) (out *ApiSSEOut, err error) {
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
