package model

import (
	"fmt"

	"golang.org/x/sync/singleflight"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"

	"github.com/xavierror/gowheel/logs"

	"github.com/xavierror/gin-api/conf"
)

var (
	DB      *gorm.DB
	Single *singleflight.Group
)

// Setup initializes the database instance
func Setup() {
	var err error

	dbOpt := "charset=utf8mb4&parseTime=True&loc=Asia%2FShanghai"

	// 前端业务
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?%s",
		conf.Database.User,
		conf.Database.Password,
		conf.Database.Host,
		conf.Database.Name,
		dbOpt,
	)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{NamingStrategy: schema.NamingStrategy{SingularTable: true}})
	if err != nil {
		logs.Fatal(fmt.Sprintf("database %s-%s setup err: %s", conf.Database.Host, conf.Database.Name, err.Error()))
	} else {
		fmt.Printf("database %s-%s setup success \n", conf.Database.Host, conf.Database.Name)
	}

	return
}
