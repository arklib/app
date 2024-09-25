package base

import (
	"demo/app/model"

	"github.com/arklib/ark/event"
)

type Events struct {
	UserCreate *event.Event[model.User]
}

func (base *Base) defineEvents() {
	base.Events = &Events{
		UserCreate: event.New[model.User](),
	}
}
