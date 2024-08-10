package config

import (
	"context"
	"os"

	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

var RDB *redis.Client

var Ctx = context.Background()

func InitDB() {
	dsn := os.Getenv("dsn")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	DB = db
}

func InitRedis() {
	RDB = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("redis_addr"),
		Password: os.Getenv("redis_password"),
		DB:       0,
	})
}
