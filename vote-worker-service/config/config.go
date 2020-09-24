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
	User      string `env:"RABBITMQ_USER,default=guest"`
	Password  string `env:"RABBITMQ_PASSWORD,default=guest"`
	QueueName string `env:"RABBITMQ_QUEUE,default=votes"`
}

// PostgresConfig stores Postgres configuration
type PostgresConfig struct {
	Hostname string `env:"POSTGRES_HOSTNAME,default=localhost"`
	Port     uint16 `env:"POSTGRES_PORT,default=5432"`
	User     string `env:"POSTGRES_USER,default=admin"`
	Password string `env:"POSTGRES_PASSWORD,default=admin"`
	Database string `env:"POSTGRES_DB,default=voting"`
	Tables   PostgresTablesConfig
}

// PostgresTablesConfig stores Postgres tables
type PostgresTablesConfig struct {
	Votes   string `env:"POSTGRES_TABLE_VOTES,default=votes"`
	Results string `env:"POSTGRES_TABLES_RESULTS,default=results"`
}

// GetConfig loads and returns application configuration
func GetConfig() (Config, error) {
	var cfg Config
	err := envdecode.StrictDecode(&cfg)
	return cfg, err
}
