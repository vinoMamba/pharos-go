package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/vinoMamba/pharos-go/internal/jwt_helper"
)

func init() {
	jwtCmd := &cobra.Command{
		Use:   "jwt_gen",
		Short: "create a jwt secret",
		Run: func(cmd *cobra.Command, args []string) {
			homeParh, err := os.UserHomeDir()
			if err != nil {
				log.Fatal(err)
			}
			keyPath := homeParh + "/.pharos"
			// check if the dir exist
			if _, err := os.Stat(keyPath); os.IsNotExist(err) {
				os.Mkdir(keyPath, 0755)
			}
			// check if the file exist
			if _, err := os.Stat(keyPath + "/hmac.key"); os.IsNotExist(err) {
				bytes, err := jwt_helper.GenerateHMACKey()
				if err != nil {
					log.Fatal(err)
				}
				ioutil.WriteFile(keyPath+"/hmac.key", bytes, 0644)
			}
			fmt.Println("jwt secret created")
		},
	}
	rootCmd.AddCommand(jwtCmd)
}
