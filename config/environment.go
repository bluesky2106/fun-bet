package config

import (
	"github.com/spf13/viper"
)

const (
	// EnvDebug : Debug mode
	EnvDebug string = "debug"
	// EnvProduction : Production mode
	EnvProduction string = "production"
)

func setDefaultEnvironment() {
	viper.SetDefault("env", EnvDebug)
}
