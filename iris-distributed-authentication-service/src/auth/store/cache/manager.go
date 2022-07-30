package cache

import (
	"context"
	"distributed/authentication/config"
	"log"

	"github.com/go-redis/redis/v8"
)

var M *Manager

type Manager struct {
	handler *redis.Client
}

func init() {
	var err error

	config.GetConfig(&c)
	M, err = New()
	if err != nil {
		log.Fatalf("[CACHE CONNECT ERROR] %v", err)
	}
}

func New() (*Manager, error) {
	m := &Manager{
		handler: redis.NewClient(
			&redis.Options{
				Addr:     c.Addr,
				Password: c.Password,
			},
		),
	}
	return m, m.handler.Ping(context.Background()).Err()
}
