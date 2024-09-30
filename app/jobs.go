package app

import "github.com/arklib/ark/job"

type Jobs struct {
	SyncUser *job.Job[uint]
}

func (app *App) initJobs() {
	driver := job.NewRedisDriver(app.Redis)
	retryDriver := job.NewDBRetryDriver(app.DB)

	app.Jobs = &Jobs{
		SyncUser: job.Define[uint](job.Config{
			Queue:       "sync_user",
			Driver:      driver,
			RetryDriver: retryDriver,
		}),
	}
}
