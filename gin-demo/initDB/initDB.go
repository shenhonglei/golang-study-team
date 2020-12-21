package initDB

import (
	"gin-demo/settings"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var Db *gorm.DB

func init() {
	var err error
	Db, err = gorm.Open(mysql.Open(settings.AppConfig.Database.Url), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	if err != nil {
		log.Panicln("数据库连接错误", err.Error())
	}
}
