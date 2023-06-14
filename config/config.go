package config

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

func LoadAppConfig() {
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	homePath, err := os.UserHomeDir()
	if err != nil {
		log.Fatalln(err)
	}
	viper.AddConfigPath(homePath + "/.pharos")
	err = viper.ReadInConfig()
	if err != nil {
		log.Fatalln(err)
	}
}
