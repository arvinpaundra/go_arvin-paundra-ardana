package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	APIPort    string `mapstructure:"APIPort"`
	JWTSecret  string `mapstructure:"JWTSecret"`
	DBAddress  string `mapstructure:"DBAddress"`
	DBUsername string `mapstructure:"DBUsername"`
	DBPassword string `mapstructure:"DBPassword"`
	DBName     string `mapstructure:"DBName"`
}

var Cfg *Config

func InitConfig() {
	cfg := &Config{}

	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("error read env: %v", err)
	}

	viper.Unmarshal(cfg)

	Cfg = cfg
}
