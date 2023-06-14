package config

import (
	"log"

	"github.com/spf13/viper"
)

func LoadAppConfig() {
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath("/root/.config/pharos")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalln(err)
	}
}
