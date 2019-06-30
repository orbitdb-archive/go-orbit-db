package cmd

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
	"github.com/orbitdb/go-orbit-db/http"
)

var createCmd = &cobra.Command{
	Use:   "create <database> <type>",
	Short: "Creates a database or opens it if it already exists.",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 2 {
			fmt.Println("Usage: orbitdb create name type")
			os.Exit(1)
		}

		params := make(map[string]string)
		params["create"] = "true"
		params["type"] = args[1]

		var r http.Request

		r.SetURL(httpHost, "/db/", args[0])
		r.SetJSONData(params)

		fmt.Println(r.Post())
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
}
