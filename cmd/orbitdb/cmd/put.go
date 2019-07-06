package cmd

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
	"github.com/orbitdb/go-orbit-db/http"
)

var putCmd = &cobra.Command{
	Use:   "put <database> <data>",
	Short: "Adds a record to a document database.",
	Long: `
The put command adds a record with <data> to a document database named
<database>.

If the database has a unique identifier and <data> contains an existing id,
the matching record will be updated.

<data> must be JSON-formatted.
`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 2 {
			fmt.Println("Usage: orbitdb put <database> <data>")
			os.Exit(1)
		}

		var r http.Request

		r.SetURL(httpHost, "/db/", args[0])
		r.Data = args[1]

		fmt.Println(r.Post())
	},
}

func init() {
	rootCmd.AddCommand(putCmd)
}
