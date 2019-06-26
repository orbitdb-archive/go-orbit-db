package cmd

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
	"github.com/orbitdb/go-orbit-db/http"
)

var addCmd = &cobra.Command{
	Use:   "add <database> <data>",
	Short: "Adds a record to the feed or eventlog database.",
	Long: `
The add command appends a new record to a database of type feed or eventlog. Add
will not work with any other database type.

When adding a record, <data> must be plain text data,

E.g.

orbitdb add feed "OrbitDB"`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 2 {
			fmt.Println("Usage: orbitdb add <database> <data>")
			os.Exit(1)
		}

		data := args[1]

		fmt.Println(http.Post(http.BuildURL(httpHost, "/db/", args[0]), data))
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
