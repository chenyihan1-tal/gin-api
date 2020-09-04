package conf

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-ini/ini"
)

var (
	App struct {
		RuntimeRootPath string

		LogSavePath string
		LogSaveName string
		LogFileExt  string
		TimeFormat  string

		RunMode      string
		HttpPort     int
		ReadTimeout  time.Duration
		WriteTimeout time.Duration

		RedLockKey string
	}
	Database struct {
		Type        string
		User        string
		Password    string
		Host        string
		Name        string
		TablePrefix string
	}
	Redis struct {
		Host        string
		Password    string
		MinIdle     int
		MaxActive   int
		IdleTimeout time.Duration
	}
)

var cfg *ini.File

// Setup initialize the configuration instance
func Setup() {
	var err error

	RunMode := os.Getenv("RunMode")
	if RunMode != "release" {
		RunMode = "dev"
	}

	cfg, err = ini.Load(fmt.Sprintf("conf/app.%s.ini", RunMode))
	if err != nil {
		log.Fatalf("conf.Setup, fail: %v", err)
	}

	mapTo("app", &App)
	mapTo("database", &Database)
	mapTo("redis", &Redis)
}

// mapTo map section
func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo %s err: %v", section, err)
	}
}
