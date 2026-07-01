package config

import (
	"os"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type Config struct {
	HTTP

	LogLevel string `env:"LOG_LEVEL" envDefault:"debug"`
}

type HTTP struct {
	Host string `env:"HTTP_HOST,required" envDefault:"localhost"`
	Port string `env:"HTTP_PORT,required" envDefault:"1111"`
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil && !os.IsNotExist(err) {
		return nil, err
	}

	var cfg Config
	cfg, err = env.ParseAs[Config]()

	return &cfg, err

}
