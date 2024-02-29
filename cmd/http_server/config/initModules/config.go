package initModules

import (
	"database/sql"
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)

var ApplicationConfig *Config

type Config struct {
	db  *sql.DB
	rdb *redis.Client
	//clc - кликхаус
	//nt - nast
}

func NewConfig() *Config {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	cfg := Config{}
	cfg.initPostgresSql()
	cfg.initRedis()
	return &cfg
}
