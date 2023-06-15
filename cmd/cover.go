package cmd

import (
	"log"
	"net/http"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

func init() {
	coverCmd := &cobra.Command{
		Use:   "cover",
		Short: "create a test coverage report",
		Run: func(cmd *cobra.Command, args []string) {
			os.Mkdir("coverage", 0755)
			if err := exec.Command("go", "test", "-coverprofile=coverage/coverage.out", "./...").Run(); err != nil {
				log.Fatalln(err)
			}
			if err := exec.Command("go", "tool", "cover", "-html=coverage/coverage.out", "-o", "coverage/index.html").Run(); err != nil {
				log.Fatalln(err)
			}
			var port string
			if len(args) > 0 {
				port = args[0]
			} else {
				port = "8888"
			}

			log.Println("coverage report available at http://localhost:" + port + "/coverage/index.html")
			if err := http.ListenAndServe(":8888", http.FileServer(http.Dir("."))); err != nil {
				log.Fatalln(err)
			}
		},
	}
	rootCmd.AddCommand(coverCmd)
}
