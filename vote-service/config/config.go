package config

import (
	"time"

	"github.com/joeshaw/envdecode"
)

// Config stores complete application configuration
type Config struct {
	HTTP       HTTPConfig
	Rabbit     RabbitConfig
	SurveyGrpc SurveyGrcpConfig
	Postgres   PostgresConfig
}

// HTTPConfig stores HTTP configuration
type HTTPConfig struct {
	Hostname     string        `env:"HTTP_HOSTNAME"`
	Port         uint16        `env:"HTTP_PORT,default=8082"`
	ReadTimeout  time.Duration `env:"HTTP_READ_TIMEOUT,default=5s"`
	WriteTimeout time.Duration `env:"HTTP_WRITE_TIMEOUT,default=10s"`
	IdleTimeout  time.Duration `env:"HTTP_IDLE_TIMEOUT,default=2m"`
}

// RabbitConfig stores RabbitMQ configuration
type RabbitConfig struct {
	Hostname  string `env:"RABBITMQ_HOSTNAME,default=localhost"`
	Port      uint16 `env:"RABBITMQ_PORT,default=5672"`
	Username  string `env:"RABBITMQ_USER,default=guest"`
	Password  string `env:"RABBITMQ_PASSWORD,default=guest"`
	QueueName string `env:"RABBITMQ_QUEUE,default=votes"`
}

// SurveyGrcpConfig stores configuration to connect to the survey gRPC service
type SurveyGrcpConfig struct {
	Hostname string `env:"SURVEY_GRPC_HOSTNAME,default=localhost"`
	Port     uint16 `env:"SURVEY_GRPC_PORT,default=9000"`
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
	Results string `env:"POSTGRES_TABLES_RESULTS,default=results"`
}

// GetConfig loads and returns application configuration
func GetConfig() (Config, error) {
	var cfg Config
	err := envdecode.StrictDecode(&cfg)
	return cfg, err
}
