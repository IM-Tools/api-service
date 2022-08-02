package model

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"im-services/internal/config"
)

var DB *gorm.DB

type BaseModel struct {
	ID int64
}

func InitDb() *gorm.DB {
	var (
		host     = config.Conf.Mysql.Host
		port     = config.Conf.Mysql.Port
		database = config.Conf.Mysql.Database
		username = config.Conf.Mysql.Username
		password = config.Conf.Mysql.Password
		charset  = config.Conf.Mysql.Charset

		err error
	)

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s",
		username, password, host, port, database, charset)

	DB, err = gorm.Open(mysql.New(mysql.Config{
		DSN: dsn,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		fmt.Println("Mysql 连接异常: ")
		panic(err)
	}

	return DB
}
