package infrastructure

import (
	"flag"
	"fmt"

	"github.com/spf13/viper"
)

type (
	Config struct {
		Server   ServerConfig   `mapstructure:"server"`
		Database DatabaseConfig `mapstructure:"database"`
	}

	ServerConfig struct {
		Port string `mapstructure:"port"`
		Env  string `mapstructure:"env"`
	}

	DatabaseConfig struct {
		Host     string `mapstructure:"host"`
		Username string `mapstructure:"username"`
		Password string `mapstructure:"password"`
		Port     string `mapstructure:"port"`
	}
)

const (
	defaultConfigFilename = "config.yml"
	defaultConfigPath     = "../etc"
	defaultConfigType     = "yaml"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config", defaultConfigPath, "path to config file")

}

func NewConfig() *Config {
	fmt.Println("Read Configuration...")

	viper.SetConfigName(defaultConfigFilename)
	viper.SetConfigType(defaultConfigType)
	viper.AddConfigPath(configPath)
	viper.AutomaticEnv()

	var config *Config

	if err := viper.ReadInConfig(); err != nil {
		panic("Error reading config: " + err.Error())
	}

	err := viper.Unmarshal(&config)
	if err != nil {
		panic("Error decode config: " + err.Error())
	}

	return config
}
