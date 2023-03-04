package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"sync"
)

type Connection struct {
	Db *gorm.DB
}

var instance *Connection
var once sync.Once

func GetDbConnection() *Connection {
	once.Do(func() {
		//temp, err := gorm.Open(postgres.Open(conf.DB.Info), &gorm.Config{
		temp, err := gorm.Open(postgres.Open("host=whisky_db_1.0 user=root password=root dbname=way port=5432 sslmode=disable TimeZone=Asia/Seoul"), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})
		instance = &Connection{temp}
		if err != nil {
			err.Error()
			panic("Db 연결에 실패하였습니다.")
		}
	})
	return instance
}

func GetDb() *Connection {
	return instance
}
