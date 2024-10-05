package app

import (
	"demo/app/user/model"
	"github.com/arklib/ark/queue"
)

type Queues struct {
	UserCreate *queue.Queue[model.User]
}

func (app *App) initQueues() {
	// brokers := []string{"localhost:9092"}
	// kafkaDriver := queue.NewKafkaDriver(brokers)
	driver := queue.NewRedisDriver(app.Redis)
	retryDriver := queue.NewDBRetryDriver(app.DB)

	queues := new(Queues)
	queues.UserCreate = queue.Define[model.User](queue.Config{
		Name:        "user_create",
		Driver:      driver,
		RetryDriver: retryDriver,
	})

	app.Queues = queues
}
