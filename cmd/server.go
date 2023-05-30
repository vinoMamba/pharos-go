package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/vinoMamba/pharos-go/internal/router"
)

func init() {
	var serverCmd = &cobra.Command{
		Use:   "server",
		Short: "Start the Pharos server",
		Run: func(cmd *cobra.Command, args []string) {
			r := router.New()
			if err := r.Run(":3000"); err != nil {
				log.Println(err)
			}
		},
	}
	rootCmd.AddCommand(serverCmd)
}
