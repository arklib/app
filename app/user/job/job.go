package job

import "demo/app"

type Job struct {
	*app.App
}

func New(app *app.App) *Job {
	return &Job{
		app,
	}
}
