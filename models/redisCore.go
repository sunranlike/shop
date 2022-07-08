package models

//https://gorm.io/zh_CN/docs/connecting_to_the_database.html
import (
	"context"
	"github.com/go-redis/redis/v8"
)

var (
	RedisDb *redis.Client
)

func init() {
	var ctx = context.Background()
	RedisDb = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})
	_, err := RedisDb.Ping(ctx).Result()
	if err != nil {
		println(err)
	}
}
