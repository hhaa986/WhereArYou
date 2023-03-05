package db

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"sync"
)

type Connection struct {
	Db *gorm.DB
}

var instance *Connection
var once sync.Once

func GetDbConnection() *Connection {
	once.Do(func() {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
		dsn := "host=" + os.Getenv("DB_HOST") +
			" user=" + os.Getenv("DB_USER") +
			" password=" + os.Getenv("DB_PWD") +
			" dbname=" + os.Getenv("DB_NAME") +
			" port=" + os.Getenv("DB_PORT") +
			" sslmode=disable TimeZone=Asia/Seoul"
		fmt.Println("dsn:" + dsn)
		temp, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
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
