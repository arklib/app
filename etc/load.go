package etc

import (
	"demo/app"
	"demo/app/test"
	"demo/app/user"
	"github.com/arklib/ark"
	"github.com/arklib/ark/config"
)

func LoadApp(configFile string) *app.App {
	conf := config.MustLoad(configFile)
	srv := ark.MustNewServer(conf)
	return app.New(srv).Use(
		test.Define,
		user.Define,
	)
}
