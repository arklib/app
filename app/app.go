package app

import (
	"demo/app/base"
	"github.com/arklib/ark"
	"github.com/arklib/ark/config"

	"demo/app/fn/test"
	"demo/app/fn/user"
)

func Run() {
	c := config.MustLoad("./config.toml")
	srv := ark.MustNewServer(c)

	// base init
	base.New(srv).Use(
		test.Setup,
		user.Setup,
	).Init()

	srv.Run()
}
