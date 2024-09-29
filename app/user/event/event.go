package event

import (
	"demo/app"
)

type Event struct {
	*app.App
}

func New(app *app.App) *Event {
	return &Event{
		app,
	}
}
