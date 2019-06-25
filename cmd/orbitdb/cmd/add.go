package cmd

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
	"github.com/orbitdb/go-orbit-db/http"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a record to the feed or eventlog database.",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 2 {
			fmt.Println("Usage: orbitdb add name data")
			os.Exit(1)
		}

		data := args[1]

		fmt.Println(http.Post(http.BuildURL(httpHost, "/db/", args[0]), data))
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
