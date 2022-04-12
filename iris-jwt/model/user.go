package model

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

type User struct {
	Id         string `json:"id"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
	Permission int    `json:"permission"`
}

var (
	DB      *sql.DB
	RedisDB *redis.Client
)

func openDB(username, password, addr, dbname string) *sql.DB {
	config := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True", username, password, addr, dbname)
	db, err := sql.Open("mysql", config)
	if err != nil {
		log.Fatalln("failed to open DB, err: ", err)
	}
	return db
}

func InitDB() {
	DB = openDB(viper.GetString("db.username"), viper.GetString("db.password"), viper.GetString("db.addr"), viper.GetString("db.name"))
}

func openRedisDB(addr, password string, db int) *redis.Client {
	redisDB := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	_, err := redisDB.Ping().Result()
	if err != nil {
		log.Fatalln("failed to open Redis, err: ", err)
	}

	return redisDB
}

func InitRedis() {
	RedisDB = openRedisDB(viper.GetString("redis.addr"), viper.GetString("redis.password"), viper.GetInt("redis.db_num"))
}

func SelectUserInfoById(id int) User {
	var user User
	row := DB.QueryRow("select id, username, password, email, phone, permission from user where id = ?", id)
	err := row.Scan(&user.Id, &user.Username, &user.Password, &user.Email, &user.Phone, &user.Permission)
	if err != nil {
		log.Println("get the user information err: ", err)
	}
	return user
}

func SelectUserInformationByUsername(username string) User {
	var user User
	row := DB.QueryRow("select id, username, password, email, phone, permission from user where username = ?", username)
	err := row.Scan(&user.Id, &user.Username, &user.Password, &user.Email, &user.Phone, &user.Permission)
	if err != nil {
		log.Println("get the user information err: ", err)
	}
	return user
}

func SetValueWithExpire(key string, value interface{}, expiration time.Duration) error {
	return RedisDB.Set(key, value, expiration).Err()
}

func SetValue(key string, value interface{}) error {
	return RedisDB.Set(key, value, 0).Err()
}

func GetValue(key string) (string, error) {
	return RedisDB.Get(key).Result()
}
