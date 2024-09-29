package service

import (
	"demo/app"
)

type Service struct {
	*app.App
}

func New(app *app.App) *Service {
	return &Service{
		app,
	}
}
