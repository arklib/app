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
		"user.print_create",
	).Notify(app.Queues.UserCreate.Push)

	app.Hooks = hooks
}
