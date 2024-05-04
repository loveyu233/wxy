package dao

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host   = "127.0.0.1"
	port   = 5432
	user   = "root"
	pwd    = "pgsql"
	dbname = "topicselect"
)

var DB *gorm.DB

func PgsqlInit() {
	dsn := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable TimeZone=Asia/Shanghai", host, port, user, dbname, pwd)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		// 处理错误
		panic("failed to connect database")
	}
	logrus.Info("pgsql connection success")
	DB = db
}
