package config

import (
	"github.com/pawplace/ledserver/server"
)

// Config stores application configuration.
type Config struct {
	Server server.Config `yaml:"server"`
}
