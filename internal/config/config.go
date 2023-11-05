package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	Port        string
	ReleaseMode bool
	InputFile   string
	LogLevel    string
	LogFile     string
}

func Load() *Config {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	viper.SetDefault("log_level", "debug")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	config := &Config{
		Port:        viper.GetString("port"),
		ReleaseMode: viper.GetBool("release_mode"),
		InputFile:   viper.GetString("input_file"),
		LogLevel:    viper.GetString("log_level"),
		LogFile:     viper.GetString("log_file"),
	}

	return config
}
