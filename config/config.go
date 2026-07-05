package config

import (
	"os"
	"time"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type Config struct {
	HTTP
	S3

	LogLevel string `env:"LOG_LEVEL" envDefault:"debug"`
}

type HTTP struct {
	Host        string        `env:"HTTP_HOST,required" envDefault:"localhost"`
	Port        string        `env:"HTTP_PORT,required" envDefault:"1111"`
	ReadTimeout time.Duration `env:"HTTP_READ_TIMEOUT"  envDefault:"10s"`
}

type S3 struct {
	Region     string `env:"S3_REGION" envDefault:"us-east-1"`
	Endpoint   string `env:"S3_ENDPOINT"`
	AccessKey  string `env:"S3_ACCESS_KEY"`
	SecretKey  string `env:"S3_SECRET_KEY"`
	BucketName string `env:"S3_BUCKET_NAME"`
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
