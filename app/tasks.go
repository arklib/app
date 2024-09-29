package app

import "github.com/arklib/ark/task"

type Tasks struct {
	SyncUser task.Task[uint]
}

func (app *App) initTasks() {
	driver := task.NewRedisDriver(app.GetRedis())

	app.Tasks = &Tasks{
		SyncUser: task.Define[uint](task.Config{
			Driver:  driver,
			Queue:   "sync_user",
			Timeout: 60,
		}),
	}
}
