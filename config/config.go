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
	Postgres

	LogLevel string `env:"LOG_LEVEL" envDefault:"debug"`
}

type HTTP struct {
	Host               string        `env:"HTTP_HOST,required"        envDefault:"localhost"`
	Port               string        `env:"HTTP_PORT,required"        envDefault:"1111"`
	ReadTimeout        time.Duration `env:"HTTP_READ_TIMEOUT"         envDefault:"10s"`
	HealthCheckTimeout time.Duration `env:"HTTP_HEALTH_CHECK_TIMEOUT" envDefault:"3s"`
	MaxUploadMB        int64         `env:"HTTP_MAX_UPLOAD_MB"        envDefault:"10"`
}

type S3 struct {
	Region     string `env:"S3_REGION"      envDefault:"us-east-1"`
	Endpoint   string `env:"S3_ENDPOINT"`
	AccessKey  string `env:"S3_ACCESS_KEY"`
	SecretKey  string `env:"S3_SECRET_KEY"`
	BucketName string `env:"S3_BUCKET_NAME"`

	LinkTTL time.Duration `env:"S3_LINK_TTL" envDefault:"15m"`
}

type Postgres struct {
	Host     string `env:"POSTGRES_HOST,required"`
	Port     string `env:"POSTGRES_PORT,required"`
	User     string `env:"POSTGRES_USER,required"`
	Password string `env:"POSTGRES_PASSWORD,required"`
	Database string `env:"POSTGRES_DB,required"`
	SSLMode  string `env:"POSTGRES_SSLMODE"           envDefault:"disable"`

	MaxOpenConns    int           `env:"POSTGRES_MAX_OPEN_CONNS"     envDefault:"25"`
	MaxIdleConns    int           `env:"POSTGRES_MAX_IDLE_CONNS"     envDefault:"25"`
	ConnMaxLifetime time.Duration `env:"POSTGRES_CONN_MAX_LIFETIME"  envDefault:"30m"`
	ConnMaxIdleTime time.Duration `env:"POSTGRES_CONN_MAX_IDLE_TIME" envDefault:"5m"`
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
