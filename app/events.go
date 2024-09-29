package app

import (
	"demo/app/user/model"

	"github.com/arklib/ark/event"
)

type Events struct {
	UserCreate *event.Event[model.User]
}

func (app *App) initEvents() {
	app.Events = &Events{
		UserCreate: event.New[model.User](),
	}
}
