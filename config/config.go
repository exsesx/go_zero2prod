package config

import (
	"github.com/spf13/viper"
	"log"
)

var config *viper.Viper

// Init is an exported method that takes the environment starts the viper
// (external lib) and returns the configuration struct.
func Init() {
	var err error

	config = viper.New()

	config.SetConfigFile(".env")
	config.AddConfigPath(".")

	err = config.ReadInConfig()

	if err != nil {
		log.Fatal("error on parsing configuration file", err)
	}
}

func GetConfig() *viper.Viper {
	return config
}
