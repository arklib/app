package user

import (
	"github.com/spf13/cobra"

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

	// add command
	app.AddCommand(&cobra.Command{
		Use: "user:sync_user_form_erp",
		Run: func(*cobra.Command, []string) {
			_ = userSvc.SyncUserFormERP()
		},
	})

	// add hook
	app.Hooks.UserCreateAfter.Add("user_create_print", userSvc.UserCreatePrint)

	// add queue task
	app.Queues.UserCreate.AddTask("send_user_create_mail", userSvc.SendUserCreateMail,
		queue.TaskConfig{MaxRetry: 1, RetryInterval: 15},
	)
	app.Queues.UserCreate.AddTask("sync_user_to_redis", userSvc.SyncUserToRedis,
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
