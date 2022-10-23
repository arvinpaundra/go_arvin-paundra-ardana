package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	APIPort    string
	DBAddress  string
	DBUsername string
	DBPassword string
	JWTSecret  string
}

var Cfg *Config

func InitConfig() {
	cfg := &Config{}

	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(err)
	}

	viper.Unmarshal(cfg)

	Cfg = cfg
}
