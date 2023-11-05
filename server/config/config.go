package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

type Config struct {
	Database struct {
		Type     string `mapstructure:"type"`
		FilePath string `mapstructure:"file_path"`
	} `mapstructure:"database"`

	Server struct {
		Port    int    `mapstructure:"port"`
		BaseURL string `mapstructure:"base_url"`
	} `mapstructure:"server"`

	Logging struct {
		Level    string `mapstructure:"level"`
		FilePath string `mapstructure:"file_path"`
	} `mapstructure:"logging"`
}

func LoadConfigData() (Config, error) {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetConfigType("yml")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("Config file not found.")
			os.Exit(1)
		} else {
			fmt.Println(err)
			os.Exit(1)
		}
	}

	var c Config
	err := viper.Unmarshal(&c)
	return c, err
}
