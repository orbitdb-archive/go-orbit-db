package cmd

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
	"github.com/orbitdb/go-orbit-db/http"
)

var dropCmd = &cobra.Command{
	Use:   "drop <database>",
	Short: "Drop a database.",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Usage: orbitdb drop <database>")
			os.Exit(1)
		}

		var r http.Request

		r.SetURL(httpHost, "/db/", args[0])

		fmt.Println(r.Delete())
	},
}

func init() {
	rootCmd.AddCommand(dropCmd)
}
