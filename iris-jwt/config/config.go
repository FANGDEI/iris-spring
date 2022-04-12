package config

import (
	"github.com/spf13/viper"
)

func Init() error {
	viper.AddConfigPath("./conf")
	viper.SetConfigName("config")
	viper.SetConfigType("yml")

	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	return nil
}
