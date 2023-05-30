package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	var serverCmd = &cobra.Command{
		Use:   "server",
		Short: "Start the Pharos server",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Starting Pharos server...")
			fmt.Println(args)
		},
	}
	rootCmd.AddCommand(serverCmd)
}
