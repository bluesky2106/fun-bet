package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// Config : configurations
type Config struct {
	Host string
	Port string
	Env  Environment
	MySQL
}

// ParseConfig : parse configurations from global env and json file
func ParseConfig(file, path string) *Config {
	// Set default variables
	setDefaultVariables()

	// Enable VIPER to read Environment Variables
	readEnvironmentVariables()

	// Parse configurations from JSON file
	readJSONFile(file, path)

	var conf Config
	err := viper.Unmarshal(&conf)
	if err != nil {
		fmt.Printf("Unable to decode into struct, %v\n", err)
	}

	return &conf
}

func setDefaultServer() {
	viper.SetDefault("host", "localhost")
	viper.SetDefault("port", "8080")
}

func setDefaultVariables() {
	setDefaultServer()
	setDefaultEnvironment()
	setDefaultMySQL()
}

func readEnvironmentVariables() {
	viper.AutomaticEnv()
}

func readJSONFile(file, path string) {
	fileName, fileType := getFileNameAndType(file)

	// Set the file name & type of the configurations file
	viper.SetConfigName(fileName)
	viper.SetConfigType(fileType)

	// Set the path to look for the configurations file
	viper.AddConfigPath(path)

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s\n", err)
	}
}

// Print configurations for checking
func (conf *Config) Print() {
	fmt.Println("---------------------- Server configurations ----------------------")
	fmt.Println("host: \t\t\t", conf.Host)
	fmt.Println("port: \t\t\t", conf.Port)
	conf.printMySQLConfig()
}
