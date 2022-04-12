package config

import (
	"io/ioutil"
	"log"

	jsoniter "github.com/json-iterator/go"
)

var Configuration = &configuration{}

func init() {
	// 指定配置文件
	b, err := ioutil.ReadFile("conf/config.json")
	if err != nil {
		log.Fatalln("failed to load the configuration")
	}
	// 配置文件加载到Configuration中
	err = jsoniter.Unmarshal(b, Configuration)
	if err != nil {
		log.Fatalln("failed to load the configuration")
	}
}

type configuration struct {
	Port  string  `json:"port"`
	Db    db      `json:"db"`
	Redis redisDB `json:"redis"`
}

type db struct {
	DBName   string `json:"dbname"`
	Addr     string `json:"addr"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type redisDB struct {
	Addr     string `json:"addr"`
	Password string `json:"password"`
	DBnum    int    `json:"db_num"`
}
