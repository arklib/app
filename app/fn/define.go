package fn

import (
	"demo/app/base"
	"demo/app/fn/test"
	"demo/app/fn/user"
)

func Define(base *base.Base) {
	router := base.HttpServer.Group("api")

	base.AddFn("test", test.New(base, router))
	base.AddFn("user", user.New(base, router))
}
