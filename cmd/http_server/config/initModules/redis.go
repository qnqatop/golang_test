package initModules

import (
	"fmt"
	"github.com/redis/go-redis/v9"
)

func (cfg *Config) initRedis() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("localhost:6379"),
		Password: "",
		DB:       0,
	})
	cfg.rdb = rdb
}

func (cfg *Config) GetRDb() *redis.Client {
	return cfg.rdb
}
