package database

import (
	"database/sql"
	"fmt"
	"golang-jwt/config"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var mysqlDB *sql.DB

func init() {
	config := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True",
		config.Configuration.Db.Username, config.Configuration.Db.Password, config.Configuration.Db.Addr, config.Configuration.Db.DBName)
	var err error
	mysqlDB, err = sql.Open("mysql", config)
	if err != nil {
		log.Fatalln("failed to connect to database")
	}
}

func GetMySQLDB() *sql.DB {
	return mysqlDB
}
