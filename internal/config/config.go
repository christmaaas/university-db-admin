package config

import (
	"log"

	"github.com/ilyakaznacheev/cleanenv"
)

type DatabaseConfig struct {
	Host     string `env:"DB_HOST"`
	Port     string `env:"DB_PORT"`
	Name     string `env:"DB_NAME"`
	User     string `env:"DB_USER"`
	Password string `env:"DB_PASS"`
}

type Config struct {
	DB DatabaseConfig
}

var cfg *Config = &Config{}

func LoadConfig() *Config {
	log.Println("reading database config")
	if err := cleanenv.ReadConfig(".env", &cfg.DB); err != nil {
		log.Fatal("cant get database config: ", err)
	}
	return cfg
}
