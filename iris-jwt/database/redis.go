package database

import (
	"iris-jwt/config"
	"log"

	"github.com/go-redis/redis"
)

var redisDB *redis.Client

func init() {
	redisDB = redis.NewClient(&redis.Options{
		Addr:     config.Configuration.Redis.Addr,
		Password: config.Configuration.Redis.Password,
		DB:       config.Configuration.Redis.DBnum,
	})

	_, err := redisDB.Ping().Result()
	if err != nil {
		log.Fatalln("failed to connect to redis")
	}
}

func GetRedisDB() *redis.Client {
	return redisDB
}
