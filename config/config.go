package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	MYSQL_USER     string
	MYSQL_PASSWORD string
	MYSQL_DATABASE string
	MYSQL_HOST     string
	MYSQL_PORT     int

	HOST string
	PORT int
}

func LoadConfig(path string) *Config {
	viper.SetConfigFile(path)
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err.Error())
	}

	var cfg Config
	err = viper.Unmarshal(&cfg)
	if err != nil {
		log.Fatal(err.Error())
	}

	return &cfg
}
