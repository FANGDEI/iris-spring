package config

import (
	"distributed/authentication/center"
	"encoding/json"
	"log"
)

type Config interface {
	Key() string
}

func GetConfig(c Config) {
	res, err := center.GetValue(c.Key())
	if err != nil {
		log.Fatalf("[LOAD CONFIG ERROR] %v", err)
	}

	err = json.Unmarshal(res, &c)
	if err != nil {
		log.Fatalf("[UNMARSHAL ERROR] %v", err)
	}
}
