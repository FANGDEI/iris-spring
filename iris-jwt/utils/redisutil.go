package utils

import (
	"iris-jwt/database"
	"time"

	"github.com/go-redis/redis"
)

var RedisDB *redis.Client = database.GetRedisDB()

func SetValueWithExpire(key string, value interface{}, expiration time.Duration) error {
	return RedisDB.Set(key, value, expiration).Err()
}

func SetValue(key string, value interface{}) error {
	return RedisDB.Set(key, value, 0).Err()
}

func GetValue(key string) (string, error) {
	return RedisDB.Get(key).Result()
}

func DeleteValue(key string) error {
	return RedisDB.Del(key).Err()
}
