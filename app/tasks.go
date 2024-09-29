package app

import "github.com/arklib/ark/task"

type Tasks struct {
	SyncUser task.Task[uint]
}

func (app *App) initTasks() {
	// 传递未初始化客户端
	// driver := task.NewRedisDriver(app.UseRedis)

	driver := task.NewRedisDriver(app.UseRedis())
	app.Tasks = &Tasks{
		SyncUser: task.New[uint](driver, "sync_user", 60),
	}
}
