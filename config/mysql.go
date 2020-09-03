package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// MySQL configurations
type MySQL struct {
	DBName   string `json: "dbName"`
	Username string `json: "username"`
	Password string `json: "password"`
	Host     string `json: "host"`
	Port     string `json: "port"`
}

func setDefaultMySQL() {
	viper.SetDefault("mysql.dbName", "test_db")
	viper.SetDefault("mysql.username", "root")
	viper.SetDefault("mysql.password", "")
	viper.SetDefault("mysql.host", "localhost")
	viper.SetDefault("mysql.port", "3306")
}

func (conf *Config) printMySQLConfig() {
	fmt.Println("MySQL DB name:\t\t", conf.MySQL.DBName)
	fmt.Println("MySQL User:\t\t", conf.MySQL.Username)
	fmt.Println("MySQL Pass:\t\t", conf.MySQL.Password)
	fmt.Println("MySQL Host:\t\t", conf.MySQL.Host)
	fmt.Println("MySQL Port:\t\t", conf.MySQL.Port)
}
