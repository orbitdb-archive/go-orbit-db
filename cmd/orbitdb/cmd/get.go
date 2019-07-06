package cmd

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
	"github.com/orbitdb/go-orbit-db/http"
)

var getCmd = &cobra.Command{
	Use:   "get <database> [<id>]",
	Short: "Gets a record from the database.",
	Long: `
The get command gets a record identified by <id> from the database named
<database>.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Usage: orbitdb get <database> [<id>]")
			os.Exit(1)
		}

		var r http.Request

		r.SetURL(httpHost, "/db/", args[0])

		// if no id is supplied, assume counter is being retrieved.
		if len(args) == 1 {
			r.SetURL(r.Url, "/value")
		} else {
			r.SetURL(r.Url, "/", args[1])
		}

		fmt.Println(r.Get())
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}
