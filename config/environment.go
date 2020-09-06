package config

import (
	"github.com/spf13/viper"
)

const (
	// EnvDevelopment : Debug mode
	EnvDevelopment string = "debug"
	// EnvProduction : Production mode
	EnvProduction string = "production"
)

func setDefaultEnvironment() {
	viper.SetDefault("env", EnvDevelopment)
}
