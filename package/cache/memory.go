package cache

import (
	"time"

	"github.com/patrickmn/go-cache"
	"github.com/xaviercry/gopkg/logs"
)

var m *cache.Cache

func memory() {
	m = cache.New(30*time.Minute, 10*time.Minute)

	m.Set("test",10086,0)
	_ , ok := m.Get("test")
	if !ok {
		logs.Error("memory cache setup fail")
	}
}
