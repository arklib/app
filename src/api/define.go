package api

import (
	. "app/src"
	"app/src/api/shop"
	"app/src/api/test"
	"app/src/api/top"
)

func Define(app *App) {
	shop.New(app)
	test.New(app)
	top.New(app)
}
