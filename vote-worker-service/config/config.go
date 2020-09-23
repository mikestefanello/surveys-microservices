package config

import (
	"github.com/joeshaw/envdecode"
)

// Config stores complete application configuration
type Config struct {
	Rabbit   RabbitConfig
	Postgres PostgresConfig
}

// RabbitConfig stores RabbitMQ configuration
type RabbitConfig struct {
	Hostname  string `env:"RABBITMQ_HOSTNAME,default=localhost"`
	Port      uint16 `env:"RABBITMQ_PORT,default=5672"`
	Username  string `env:"RABBITMQ_USERNAME,default=guest"`
	Password  string `env:"RABBITMQ_PASSWORD,default=guest"`
	QueueName string `env:"RABBITMQ_QUEUE,default=votes"`
}

// PostgresConfig stores PostgresDB configuration
type PostgresConfig struct {
	// TODO
}

// GetConfig loads and returns application configuration
func GetConfig() (Config, error) {
	var cfg Config
	err := envdecode.StrictDecode(&cfg)
	return cfg, err
}
