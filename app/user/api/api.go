package api

import (
	"demo/app"
	testapi "demo/app/test/api"
	"demo/app/user/service"
)

type Api struct {
	*app.App
	testApi *testapi.Api
	userSvc *service.Service
}

func New(app *app.App) *Api {
	return &Api{
		app,
		testapi.New(app),
		service.New(app),
	}
}
