package api

import (
	"demo/app"
)

type Api struct {
	*app.App
}

func New(app *app.App) *Api {
	return &Api{
		app,
	}
}
