package api

import (
	"demo/app"
	testapi "demo/app/test/api"
)

type Api struct {
	*app.App
	testApi *testapi.Api
}

func New(app *app.App) *Api {
	return &Api{
		app,
		app.GetApi("test").(*testapi.Api),
	}
}
