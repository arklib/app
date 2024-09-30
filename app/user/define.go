package user

import (
	"demo/app"
	"demo/app/user/api"
	"demo/app/user/event"
	"demo/app/user/job"
	"demo/app/user/model"
	"demo/app/user/service"
	"demo/app/user/task"
	. "github.com/arklib/ark"
)

func Define(app *app.App) {
	// add model
	app.AddModel("user", new(model.User))
	app.AddModel("user_address", new(model.UserAddress))

	// add service
	userService := service.New(app)
	app.AddService("user", userService)

	// add event
	userEvent := event.New(app)
	app.Events.UserCreate.Use(userEvent.SendUserCreateSMS)

	// add job
	userJob := job.New(app)
	app.Jobs.SyncUser.Use(userJob.SyncUser)

	// register task
	userTask := task.New(app)
	app.Task.Define("user:SyncUserFormERP", userTask.SyncUserFormERP)

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

	router = app.HttpServer.Group("api/user")
	router.AddRoutes(HttpRoutes{
		{
			Path:    "get",
			Handler: ApiHandler[api.GetIn, api.GetOut](apiFn.Get),
		},
		{
			Path:    "search",
			Handler: ApiHandler[api.SearchIn, api.SearchOut](apiFn.Search),
		},
	}, app.Auth.HttpMiddleware("user"))
}
