// Package config is responsible for storing configuration data of the application.
// Config struct is used to load data from environment variables.
// `envconfig` is a specific tag for https://github.com/kelseyhightower/envconfig package.
package config

import "time"

type Config struct {
	LogLevel           string        `envconfig:"LOG_LEVEL" default:"debug"`
	ServerPort         int           `envconfig:"SERVER_PORT" default:"8080"`
	ServerReadTimeout  time.Duration `envconfig:"SERVER_READ_TIMEOUT" default:"15s"`
	ServerWriteTimeout time.Duration `envconfig:"SERVER_WRITE_TIMEOUT" default:"15s"`
}
