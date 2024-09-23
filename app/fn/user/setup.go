package user

import (
	"demo/app/base"
	"demo/app/fn/test"
	. "github.com/arklib/ark"
)

type Fn struct {
	*base.Base
	test *test.Fn
}

func Setup(base *base.Base) {
	// add fn
	fn := &Fn{
		base,
		base.Fn("test").(*test.Fn),
	}
	base.AddFn("user", fn)

	// add event
	base.Events.UserCreate.Add(
		fn.OnCreateSendSMS,
		fn.OnCreateSendEmail,
	)

	// add api
	router := base.HttpServer
	api := router.Group("api/user")
	api.AddRoutes(HttpRoutes{
		{Path: "login", Handler: ApiHandler[ApiLoginIn, ApiLoginOut](fn.ApiLogin)},
		{Path: "create", Handler: ApiHandler[ApiCreateIn, ApiCreateOut](fn.ApiCreate)},
	})

	authMw := base.Auth.HttpMiddleware("user")
	api.AddRoutes(HttpRoutes{
		{Path: "get", Handler: ApiHandler[ApiGetIn, ApiGetOut](fn.ApiGet)},
		{Path: "search", Handler: ApiHandler[ApiSearchIn, SearchOut](fn.ApiSearch)},
	}, authMw)
}
