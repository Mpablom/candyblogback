package config

import (
	"fmt"
	"log"
	"os"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var DB *gorm.DB

type DbConfig struct {
	Host         string
	Port         string
	Database     string
	User         string
	Password     string
	Charset      string
	MaxIdleConns int
	MaxOpenConns int
	TimeZone     string
	print_log    bool
}

func Configure(ConfigPath string, ConfigName string) DbConfig {
	viper.AddConfigPath(ConfigPath)
	viper.SetConfigName(ConfigName)
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
	if viper.GetString("default.host") == "" {
		fmt.Println("falta host en el archivo " + ConfigName)
		os.Exit(1)
	}
	if viper.GetString("default.database") == "" {
		fmt.Println("falta database en el archivo de config " + ConfigName)
		os.Exit(1)
	}
	if viper.GetString("default.user") == "" {
		fmt.Println("falta user en el archivo de config " + ConfigName)
		os.Exit(1)
	}
	if viper.GetString("default.password") == "" {
		fmt.Println("falta password en el archivo de config " + ConfigName)
		os.Exit(1)
	}

	response := DbConfig{
		viper.GetString("default.host"),
		viper.GetString("default.port"),
		viper.GetString("default.database"),
		viper.GetString("default.user"),
		viper.GetString("default.password"),
		viper.GetString("default.charset"),
		viper.GetInt("default.MaxIdleConns"),
		viper.GetInt("default.MaxOpenConns"),
		viper.GetString("default.time_zone"),
		viper.GetBool("default.sql_log"),
	}
	if response.TimeZone == "" {
		response.TimeZone = "UTC"
	}

	return response
}
