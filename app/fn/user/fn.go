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

func New(base *base.Base, router *HttpRouter) *Fn {
	fn := &Fn{
		base,
		base.GetFn("test").(*test.Fn),
	}

	// add event
	base.Events.UserCreate.Add(fn.OnCreateSendSMS)

	// add api
	userRouter := router.Group("user")
	userRouter.AddRoutes(HttpRoutes{
		{
			Path:    "login",
			Handler: ApiHandler[ApiLoginIn, ApiLoginOut](fn.ApiLogin),
		},
		{
			Path:    "create",
			Handler: ApiHandler[ApiCreateIn, ApiCreateOut](fn.ApiCreate),
		},
	})

	authMw := base.Auth.HttpMiddleware("user")
	userRouter.AddRoutes(HttpRoutes{
		{
			Path:    "get",
			Handler: ApiHandler[ApiGetIn, ApiGetOut](fn.ApiGet),
		},
		{
			Path:    "search",
			Handler: ApiHandler[ApiSearchIn, ApiSearchOut](fn.ApiSearch),
		},
	}, authMw)

	return fn
}
