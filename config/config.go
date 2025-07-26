// Package config handles application configuration loading and management.
package config

import (
	"log"

	"github.com/spf13/viper"
)

type DBConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Name     string
}

type Config struct {
	Database DBConfig
}

var AppConfig Config

func LoadConfig() {
	viper.SetConfigName("config") // config.yaml
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config") // config folder

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("❌ Failed to read config: %v", err)
	}

	err = viper.Unmarshal(&AppConfig)
	if err != nil {
		log.Fatalf("❌ Failed to unmarshal config: %v", err)
	}

	log.Println("✅ Config loaded successfully")
}
