package user

import (
	"demo/app"
	"demo/app/user/api"
	"demo/app/user/model"
	"demo/app/user/service"
	. "github.com/arklib/ark"
	"github.com/arklib/ark/queue"
)

func Define(app *app.App) {
	// add model
	app.AddModel("user", new(model.User))
	app.AddModel("user_address", new(model.UserAddress))

	// service
	userSvc := service.New(app)

	// add custom task
	app.Task.Add("user.sync_form_erp", userSvc.SyncFormERP)

	// add hook
	app.Hooks.UserCreateAfter.Add("user.print_create", userSvc.PrintCreate)

	// add queue task
	app.Queues.UserCreate.AddTask("user.send_create_mail", userSvc.SendCreateMail,
		queue.TaskConfig{MaxRetry: 1, RetryInterval: 15},
	)
	app.Queues.UserCreate.AddTask("user.sync_to_redis", userSvc.SyncToRedis,
		queue.TaskConfig{RetryInterval: 5},
	)

	// api
	userApi := api.New(app)

	// add routes
	router := app.HttpServer.Group("api/user")
	router.AddRoutes(HttpRoutes{
		{
			Path:    "login",
			Handler: ApiHandler[api.LoginIn, api.LoginOut](userApi.Login),
		},
		{
			Path:    "create",
			Handler: ApiHandler[api.CreateIn, api.CreateOut](userApi.Create),
		},
	})

	router = app.HttpServer.Group("api/user")
	router.AddRoutes(HttpRoutes{
		{
			Path:    "get",
			Handler: ApiHandler[api.GetIn, api.GetOut](userApi.Get),
		},
		{
			Path:    "search",
			Handler: ApiHandler[api.SearchIn, api.SearchOut](userApi.Search),
		},
	}, app.Auth.HttpMiddleware("user"))
}
