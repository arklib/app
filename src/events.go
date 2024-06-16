package src

import (
	"github.com/arklib/ark/emitter"

	"app/src/dal/model"
)

type Events struct {
	UserCreate emitter.Emitter[model.User]
}

func initEvents(app *App) *Events {
	return &Events{
		UserCreate: emitter.New[model.User](
			app.SMS.OnUserCreate,
			app.Email.OnUserCreate,
		),
	}
}
