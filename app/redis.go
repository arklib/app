package app

import (
	"log"

	"github.com/redis/go-redis/v9"
)

func (app *App) initRedis() {
	var c struct {
		Addrs    []string `default:":6379"`
		DB       int
		Password string
	}

	err := app.BindConfig("redis", &c)
	if err != nil {
		log.Fatalf("redis config: %v", err)
	}

	redisInst := redis.NewUniversalClient(&redis.UniversalOptions{
		Addrs:    c.Addrs,
		DB:       c.DB,
		Password: c.Password,
	})
	app.Redis = redisInst
}
