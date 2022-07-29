package local

import (
	"distributed/authentication/config"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var M *Manager

type Manager struct {
	handler *gorm.DB
}

func init() {
	var err error

	config.GetConfig(&c)
	M, err = New()
	if err != nil {
		log.Fatalf("[LOCAL CONNECT ERROR] %v", err)
	}
}

func New() (*Manager, error) {
	db, err := gorm.Open(
		mysql.New(mysql.Config{
			DSN: fmt.Sprintf("%s:%s@tcp(%s:%s)/fang?charset=utf8mb4&parseTime=True&loc=Local",
				c.Username,
				c.Password,
				c.Addr,
				c.Port,
			),
		}),
		&gorm.Config{},
	)
	return &Manager{
		handler: db,
	}, err
}
