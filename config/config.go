package config

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	App      AppConfig    `mapstructure:"app"`
	Server   ServerConfig `mapstructure:"server"`
	Database Database     `mapstructure:"database"`
}

type AppConfig struct {
	Name    string `mapstructure:"name"`
	Version string `mapstructure:"version"`
}

type ServerConfig struct {
	Port string `mapstructure:"port"`
	Host string `mapstructure:"host"`
}

type Database struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Name     string `mapstructure:"name"`
}

func LoadConfig(path string) (*Config, error) {
	viper.SetConfigName("config")
	viper.AddConfigPath(path)
	viper.SetConfigType("yaml")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error unmarshaling config, %s", err)
		return nil, err
	}
	var config Config
	err = viper.Unmarshal(&config)
	if err != nil {
		log.Fatalf("Error unmarshaling config , %s", err)
		return nil, err
	}
	return &config, nil
}
