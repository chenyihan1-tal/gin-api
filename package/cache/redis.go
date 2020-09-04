package cache

import (
	"github.com/xavierror/gin-api/conf"

	Redis "github.com/go-redis/redis"
	"github.com/xavierror/gowheel/logs"
)

var (
	R *Redis.Client
)

func redis() {
	R = Redis.NewClient(&Redis.Options{
		Addr:         conf.Redis.Host,
		Password:     conf.Redis.Password,
		IdleTimeout:  conf.Redis.IdleTimeout,
		PoolSize:     conf.Redis.MaxActive,
		MinIdleConns: conf.Redis.MinIdle,
	})

	_, err := R.Ping().Result()
	if err != nil {
		logs.Warn("redis setup err: " + err.Error())
	}
}
