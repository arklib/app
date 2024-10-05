package app

import (
	"demo/app/user/model"
	"github.com/arklib/ark/hook"
)

type Hooks struct {
	UserCreateAfter *hook.Hook[model.User]
}

func (app *App) initHooks() {
	hooks := new(Hooks)
	hooks.UserCreateAfter = hook.Define[model.User](
		"user_create_print",
	)
	hooks.UserCreateAfter.Notify(app.Queues.UserCreate.Send)

	app.Hooks = hooks
}
