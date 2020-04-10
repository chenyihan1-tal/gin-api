package cache

import (
	"github.com/xaviercry/gin-api/conf"

	Redis "github.com/go-redis/redis"
	"github.com/xaviercry/gopkg/logs"
)

var (
	r *Redis.Client
)

func redis() {
	r = Redis.NewClient(&Redis.Options{
		Addr:         conf.Redis.Host,
		Password:     conf.Redis.Password,
		IdleTimeout:  conf.Redis.IdleTimeout,
		PoolSize:     conf.Redis.MaxActive,
		MinIdleConns: conf.Redis.MinIdle,
	})

	_, err := r.Ping().Result()
	if err != nil {
		logs.Warn("redis setup err: ", err.Error())
	}
}
