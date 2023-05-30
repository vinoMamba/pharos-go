package main

import (
	"github.com/spf13/viper"
	"github.com/vinoMamba/pharos-go/cmd"
	"log"
)

func main() {
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalln(err)
	}
	cmd.Execute()
}
