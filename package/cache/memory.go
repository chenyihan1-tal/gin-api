package cache

import (
	"time"

	"github.com/patrickmn/go-cache"
	"github.com/xavierror/gowheel/logs"
)

var M *cache.Cache

func memory() {
	M = cache.New(30*time.Minute, 10*time.Minute)

	M.Set("test",10086,0)
	_ , ok := M.Get("test")
	if !ok {
		logs.Error("memory cache setup fail")
	}
}
