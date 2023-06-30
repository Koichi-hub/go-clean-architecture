package config

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	MODE string

	MYSQL_USER     string
	MYSQL_PASSWORD string
	MYSQL_DATABASE string
	MYSQL_HOST     string
	MYSQL_PORT     int

	HOST string
	PORT int
}

func LoadConfig() *Config {
	path := os.Getenv("DOTENV_PATH")

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
