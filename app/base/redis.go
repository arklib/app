package base

import (
	"log"

	"github.com/redis/go-redis/v9"
)

func (base *Base) initRedis() {
	base.Logger.Info("[app] Init redis")
	config := new(struct {
		Addrs    []string `default:":6379"`
		DB       int
		Password string
	})

	err := base.BindConfig("redis", config)
	if err != nil {
		log.Fatalf("redis config: %v", err)
	}

	redisInst := redis.NewUniversalClient(&redis.UniversalOptions{
		Addrs:    config.Addrs,
		DB:       config.DB,
		Password: config.Password,
	})
	base.Redis = redisInst
}
