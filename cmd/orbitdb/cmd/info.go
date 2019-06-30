package cmd

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
	"github.com/orbitdb/go-orbit-db/http"
)

var infoCmd = &cobra.Command{
	Use:   "info <database>",
	Short: "Lists details about the database.",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Usage: orbitdb info <database>")
			os.Exit(1)
		}

		var r http.Request

		r.SetURL(httpHost, "/db/", args[0])

		fmt.Println(r.Get())
	},
}

func init() {
	rootCmd.AddCommand(infoCmd)
}
