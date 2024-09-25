package base

import (
	"log"

	"github.com/redis/go-redis/v9"
)

func (base *Base) GetRedis() redis.UniversalClient {
	if base.Redis != nil {
		return base.Redis
	}

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
	base.Logger.Debug("[app] init redis")

	base.Redis = redisInst
	return base.Redis
}
