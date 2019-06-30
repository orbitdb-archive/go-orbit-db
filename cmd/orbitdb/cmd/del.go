package cmd

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
	"github.com/orbitdb/go-orbit-db/http"
)

var delCmd = &cobra.Command{
	Use:   "del <database> <id>",
	Short: "Deletes a record from the database.",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 2 {
			fmt.Println("Usage: orbitdb del <database> <id>")
			os.Exit(1)
		}

		var r http.Request

		r.SetURL(httpHost, "/db/", args[0], "/", args[1])

		fmt.Println(r.Delete())
	},
}

func init() {
	rootCmd.AddCommand(delCmd)
}
