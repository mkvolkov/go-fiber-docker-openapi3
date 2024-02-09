package config

import (
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
)

type Server struct {
	Host string `koanf:"host"`
	Port string `koanf:"port"`
}

type Postgres struct {
	Host     string `koanf:"host"`
	Port     string `koanf:"port"`
	Username string `koanf:"username"`
	Password string `koanf:"password"`
	DBName   string `koanf:"dbname"`
}

type Config struct {
	Srv Server `koanf:"server"`

	PgCfg Postgres `koanf:"postgres"`
}

func ReadConfig(path string) (*Config, error) {
	cfg := &Config{}

	k := koanf.New(".")
	err := k.Load(file.Provider(path), yaml.Parser())
	if err != nil {
		return nil, err
	}

	err = k.Unmarshal("", &cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
