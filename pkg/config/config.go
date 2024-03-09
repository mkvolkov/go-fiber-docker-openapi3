package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Port       string `mapstructure:"PORT"`
	DBHost     string `mapstructure:"POSTGRES_HOST"`
	DBPort     string `mapstructure:"POSTGRES_PORT"`
	DBUsername string `mapstructure:"POSTGRES_USER"`
	DBPassword string `mapstructure:"POSTGRES_PASSWORD"`
	DBName     string `mapstructure:"POSTGRES_DB"`
}

func LoadConfig() (cfg Config, err error) {
	viper.AddConfigPath(".")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()

	if err != nil {
		return
	}

	err = viper.Unmarshal(&cfg)

	return
}
