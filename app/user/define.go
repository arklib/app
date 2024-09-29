package user

import (
	"demo/app"
	"demo/app/user/api"
	"demo/app/user/event"
	"demo/app/user/model"
	"demo/app/user/service"
	"demo/app/user/task"
	. "github.com/arklib/ark"
)

func Define(app *app.App) {
	// add model
	app.AddModel("user", new(model.User))

	// add service
	app.AddService("user", service.New(app))

	// add event
	userEvent := event.New(app)
	app.Events.UserCreate.Add(userEvent.SendUserCreateSMS)

	// add task
	userTask := task.New(app)
	app.Tasks.SyncUser.With(userTask.SyncFromERP)

	// add api
	apiFn := api.New(app)
	app.AddApi("user", apiFn)

	// add routes
	router := app.HttpServer.Group("api/user")
	router.AddRoutes(HttpRoutes{
		{
			Path:    "login",
			Handler: ApiHandler[api.LoginIn, api.LoginOut](apiFn.Login),
		},
		{
			Path:    "create",
			Handler: ApiHandler[api.CreateIn, api.CreateOut](apiFn.Create),
		},
	})

	authRouter := app.HttpServer.Group("api/user")
	authRouter.WithHttpMiddleware(app.Auth.HttpMiddleware("user"))
	authRouter.AddRoutes(HttpRoutes{
		{
			Path:    "get",
			Handler: ApiHandler[api.GetIn, api.GetOut](apiFn.Get),
		},
		{
			Path:    "search",
			Handler: ApiHandler[api.SearchIn, api.SearchOut](apiFn.Search),
		},
	})
}
