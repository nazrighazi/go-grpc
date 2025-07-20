package config

import (
	"fmt"
	"strings"
	"sync"

	"github.com/spf13/viper"
)

// Config struct from config.yml
type (
	ServerDetail struct {
		Port int    `mapstructure:"port"`
		Host string `mapstructure:"host"`
	}


	LogConfig struct {
		Level     string `json:"level"`
		Filename  string `json:"filename"`
		MaxSize   int    `json:"max_size"`
		MaxBackup int    `json:"max_backup"`
		MaxAge    int    `json:"max_age"`
		Compress  bool   `json:"compress"`
	}

	DbConfig struct {
		Host     string `mapstructure:"host"`
		Port     int    `mapstructure:"port"`
		User     string `mapstructure:"user"`
		Password string `mapstructure:"password"`
		DBName   string `mapstructure:"dbname"`
		SSLMode  string `mapstructure:"sslmode"`
		TimeZone string `mapstructure:"timezone"`
	}

	Config struct {
		Server              *ServerDetail `json:"server"`
		Log                 *LogConfig    `json:"log"`
		Db                  *DbConfig     `json:"db"`
	}
)

var (
	once           sync.Once
	configInstance *Config
)

// Return config instance from config.yml
func LoadConfig() (*Config, error) {
	once.Do(func() {
		viper.SetConfigName("config")
		viper.SetConfigType("json")
		viper.AddConfigPath("./users")
		viper.AutomaticEnv()
		viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

		if err := viper.ReadInConfig(); err != nil {
			fmt.Println("Error reading config file:", err)
			return
		}

		if err := viper.Unmarshal(&configInstance); err != nil {
			fmt.Println("Error unmarshaling config:", err)
			return
		}
	})
	return configInstance, nil
}