package cmd

import (
	"github.com/spf13/cobra"
	"github.com/vinoMamba/pharos-go/internal/email"
)

func init() {
	mailCmd := &cobra.Command{
		Use:   "email",
		Short: "Send email to user",
		Run: func(cmd *cobra.Command, args []string) {
			email.Send()
		},
	}
	rootCmd.AddCommand(mailCmd)
}
