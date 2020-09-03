package config

import (
	"github.com/spf13/viper"
)

// Environment : debug or production mode
type Environment string

const (
	// Debug mode
	Debug Environment = "debug"
	// Production mode
	Production Environment = "production"
)

func setDefaultEnvironment() {
	viper.SetDefault("env", Debug)
}
