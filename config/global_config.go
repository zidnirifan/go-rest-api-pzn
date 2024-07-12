package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Port     string
	Database DBConfig
	ApiKey   string
}

type DBConfig struct {
	User     string
	Password string
	Host     string
	Port     string
	DBName   string
}

var config Config

func init() {
	viper := viper.New()
	viper.SetConfigFile(".env")
	viper.AddConfigPath("../")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
	}

	config = Config{
		Port: viper.GetString("PORT"),
		Database: DBConfig{
			User:     viper.GetString("DB_USER"),
			Password: viper.GetString("DB_PASSWORD"),
			Host:     viper.GetString("DB_HOST"),
			Port:     viper.GetString("DB_PORT"),
			DBName:   viper.GetString("DB_NAME"),
		},
		ApiKey: viper.GetString("API_KEY"),
	}
}

func GetConfig() *Config {
	return &config
}
