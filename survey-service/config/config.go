package config

import (
	"time"

	"github.com/joeshaw/envdecode"
)

// Config stores complete application configuration
type Config struct {
	HTTP       HTTPConfig
	Mongo      MongoConfig
	Repository string `env:"REPOSITORY,default=mongo"`
}

// HTTPConfig stores HTTP configuration
type HTTPConfig struct {
	Hostname     string        `env:"HTTP_HOSTNAME"`
	Port         uint16        `env:"HTTP_PORT,default=8081"`
	ReadTimeout  time.Duration `env:"HTTP_READ_TIMEOUT,default=5s"`
	WriteTimeout time.Duration `env:"HTTP_WRITE_TIMEOUT,default=10s"`
	IdleTimeout  time.Duration `env:"HTTP_IDLE_TIMEOUT,default=2m"`
}

// MongoConfig stores Mongo DB configuration
type MongoConfig struct {
	URL     string        `env:"MONGO_URL,default=mongodb://localhost:27017"`
	DB      string        `env:"MONGO_DB,default=surveys"`
	Timeout time.Duration `env:"MONGO_TIMEOUT,default=5s"`
}

// GetConfig loads and returns application configuration
func GetConfig() (Config, error) {
	var cfg Config
	err := envdecode.StrictDecode(&cfg)
	return cfg, err
}
