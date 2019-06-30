package cmd

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
	"github.com/orbitdb/go-orbit-db/http"
)

var queryCmd = &cobra.Command{
	Use:   "query <database> [params]",
	Short: "Executes a query against the database.",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Usage: orbitdb query <database> [params]")
			os.Exit(1)
		}

		var r http.Request

		if len(args) == 1 {
			r.SetURL(httpHost, "/db/", args[0], "/all")
			fmt.Println(r.Get())
		} else {
			r.SetURL(httpHost, "/db/", args[0], "/query")
			r.Data = args[1]
			fmt.Println(r.Post())
		}
	},
}

func init() {
	rootCmd.AddCommand(queryCmd)
}
