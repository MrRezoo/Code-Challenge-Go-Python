package config

import (
	"errors"
	"github.com/spf13/viper"
	"log"
	"os"
)

type Config struct {
	Server ServerConfig
	Cors   CorsConfig
	Redis  RedisConfig
}

type ServerConfig struct {
	Port    string
	runMode string
}

type CorsConfig struct {
	AllowOrigin string
}

type RedisConfig struct {
	Host     string
	Port     string
	Password string
	DB       int
}

func GetConfig() *Config {
	configPath := getConfigPath(os.Getenv("APP_ENV"))
	v, err := LoadConfig(configPath, "yaml")
	if err != nil {
		log.Fatalf("unable to load config: %v", err)
	}
	config, err := ParseConfig(v)
	if err != nil {
		log.Fatalf("unable to parse config: %v", err)
	}
	return config
}

func ParseConfig(v *viper.Viper) (*Config, error) {
	var config Config
	err := v.Unmarshal(&config)
	if err != nil {
		log.Printf("unable to parse config: %v", err)
		return nil, err
	}
	return &config, nil
}

func LoadConfig(fileName string, fileType string) (*viper.Viper, error) {
	v := viper.New()
	v.SetConfigName(fileName)
	v.SetConfigType(fileType)
	v.AddConfigPath(".")
	v.AutomaticEnv()

	err := v.ReadInConfig()
	if err != nil {
		log.Printf("unable to read config: %v", err)
		var configFileNotFoundError viper.ConfigFileNotFoundError
		if errors.As(err, &configFileNotFoundError) {
			return nil, errors.New("config file not found")
		}
		return nil, err
	}
	return v, nil
}

func getConfigPath(env string) string {
	configPaths := map[string]string{
		"development": "/config/config-development",
		"production":  "/config/config-production",
		"docker":      "/config/config-docker",
	}

	path, exists := configPaths[env]
	if !exists {
		return configPaths["development"]
	}
	return path
}
