package main

import (
	"flag"

	"github.com/arklib/ark"
	"github.com/arklib/ark/config"

	"app/src"
	"app/src/api"
	"app/src/dal"
)

type Args struct {
	update bool
	config string
}

func cmdArgs() *Args {
	args := new(Args)
	flag.StringVar(&args.config, "c", "config.toml", "config file")
	flag.BoolVar(&args.update, "update", false, "update")
	flag.Parse()
	return args
}

func main() {
	args := cmdArgs()
	c := config.MustLoad(args.config)
	srv := ark.MustNewServer(c)

	app := src.NewApp(srv)
	if args.update {
		dal.Update(app)
		return
	}
	app.Use(api.Define)

	srv.Run()
}
