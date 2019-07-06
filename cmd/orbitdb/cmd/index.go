package cmd

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
	"github.com/orbitdb/go-orbit-db/http"
)

// indexCmd represents the index command
var indexCmd = &cobra.Command{
	Use:   "index <database>",
	Short: "Gets information about the data stored in <database>.",
	Long: `
The index command lists all informatoin about data stored in the database
named <database>.

The information returned will differ depending on the data type.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Usage: orbitdb index name")
			os.Exit(1)
		}

		var r http.Request

		r.SetURL(httpHost, "/db/", args[0], "/index")

		fmt.Println(r.Get())
	},
}

func init() {
	rootCmd.AddCommand(indexCmd)
}
