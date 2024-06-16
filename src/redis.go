package src

import (
	"log"

	"github.com/redis/go-redis/v9"
)

func initRedis(app *App) redis.UniversalClient {
	app.Logger.Info("[app] init redis")
	config := new(struct {
		Addrs    []string `default:":6379"`
		DB       int
		Password string
	})

	err := app.BindConfig("redis", config)
	if err != nil {
		log.Fatalf("redis config: %v", err)
	}

	rdb := redis.NewUniversalClient(&redis.UniversalOptions{
		Addrs:    config.Addrs,
		DB:       config.DB,
		Password: config.Password,
	})
	return rdb
}
