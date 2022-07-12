package models

//https://gorm.io/zh_CN/docs/connecting_to_the_database.html
import (
	"context"
	"github.com/go-redis/redis/v8"
	"gopkg.in/ini.v1"
)

var config, _ = ini.Load("./conf/app.ini")

var (
	RedisDb *redis.Client
)

//配置全局参数 用来配置redisStore

var RedisAddr = config.Section("redis").Key("ip").String()
var RedisProt = config.Section("redis").Key("port").String()
var RedisPassword = config.Section("redis").Key("password").String()

func init() {

	var ctx = context.Background()
	RedisDb = redis.NewClient(&redis.Options{
		Addr:     RedisAddr + ":" + RedisProt,
		Password: RedisPassword,
		DB:       0,
	})
	_, err := RedisDb.Ping(ctx).Result()
	if err != nil {
		println(err)
	}
}
