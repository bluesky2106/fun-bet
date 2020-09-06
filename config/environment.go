package config

import (
	"github.com/spf13/viper"
)

const (
	// EnvDevelopment : Development mode
	EnvDevelopment string = "development"
	// EnvProduction : Production mode
	EnvProduction string = "production"
)

func setDefaultEnvironment() {
	viper.SetDefault("env", EnvDevelopment)
}
