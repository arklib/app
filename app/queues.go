package app

import (
	"demo/app/user/model"
	"github.com/arklib/ark/queue"
	"github.com/arklib/ark/queue/driver"
	"github.com/arklib/ark/queue/retry"
)

type Queues struct {
	UserCreate *queue.Queue[model.User]
}

func (app *App) initQueues() {
	// kafkaDriver := driver.NewKafkaDriver("localhost:9092")
	redisDriver := driver.NewRedisDriver(app.Redis).WithTTL(86400 * 7)
	dbRetryDriver := retry.NewDBRetryDriver(app.DB)

	queues := new(Queues)
	queues.UserCreate = queue.Define[model.User](queue.Config{
		Name:        "user_create",
		Driver:      redisDriver,
		RetryDriver: dbRetryDriver,
	})
	app.Queues = queues
}
