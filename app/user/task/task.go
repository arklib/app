package task

import "demo/app"

type Task struct {
	*app.App
}

func New(app *app.App) *Task {
	return &Task{
		app,
	}
}
