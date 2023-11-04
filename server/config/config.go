package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

type Config struct {
	//TODO: Add whole infos here.
	DB   string `mapstructure:"db"`
	Port int    `mapstructure:"port"`
}

type SQLite struct {
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
