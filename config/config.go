package config

import (
	"github.com/pawplace/ledserver/leds"
	"github.com/pawplace/ledserver/server"
)

// Config stores application configuration.
type Config struct {
	Leds   leds.Config   `yaml:"leds"`
	Server server.Config `yaml:"server"`
}
