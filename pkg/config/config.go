package config

import (
	"github.com/knadh/koanf/parsers/dotenv"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
)

type Config struct {
	Port       string `koanf:"PORT"`
	DBHost     string `koanf:"POSTGRES_HOST"`
	DBPort     string `koanf:"POSTGRES_PORT"`
	DBUsername string `koanf:"POSTGRES_USER"`
	DBPassword string `koanf:"POSTGRES_PASSWORD"`
	DBName     string `koanf:"POSTGRES_DB"`
}

func LoadConfig() (cfg Config, err error) {
	k := koanf.New(".")

	if err := k.Load(file.Provider(".env"), dotenv.Parser()); err != nil {
		return cfg, err
	}

	if err := k.Unmarshal("", &cfg); err != nil {
		return cfg, err
	}

	return cfg, nil
}
