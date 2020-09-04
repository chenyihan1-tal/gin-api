package redLock

import (
	"time"

	"github.com/xavierror/gowheel/logs"

	"github.com/xavierror/gin-api/conf"
	"github.com/xavierror/gin-api/package/cache"
)

var Worker bool

func Setup() {
	// TryLock
	if !cache.R.SetNX(conf.App.RedLockKey, 1, 10*time.Second).Val() {
		time.Sleep(3 * time.Second)
		Setup()
		return
	}

	Worker = true
	logs.Info("I'm The Worker")

	// Lock
	for {
		cache.R.Set(conf.App.RedLockKey, 1, 10*time.Second)
		time.Sleep(3 * time.Second)
	}
}
