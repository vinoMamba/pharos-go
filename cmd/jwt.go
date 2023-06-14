package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/vinoMamba/pharos-go/internal/jwt_helper"
)

func init() {
	jwtCmd := &cobra.Command{
		Use:   "jwt_gen",
		Short: "create a jwt secret",
		Run: func(cmd *cobra.Command, args []string) {
			keyPahth := viper.GetString("SECRET_PATH")
			// check if the dir exist
			if _, err := os.Stat(keyPahth); os.IsNotExist(err) {
				os.Mkdir(keyPahth, 0755)
			}
			// check if the file exist
			if _, err := os.Stat(keyPahth + "/hmac.key"); os.IsNotExist(err) {
				bytes, err := jwt_helper.GenerateHMACKey()
				if err != nil {
					log.Fatal(err)
				}
				ioutil.WriteFile(keyPahth+"/hmac.key", bytes, 0644)
			}
			fmt.Println("jwt secret created")
		},
	}
	rootCmd.AddCommand(jwtCmd)
}
