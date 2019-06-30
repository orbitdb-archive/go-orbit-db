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
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Usage: orbitdb get <database> [<id>]")
			os.Exit(1)
		}

		url := http.BuildURL(httpHost, "/db/", args[0])

		// if no id is supplied, assume counter is being retrieved.
		if len(args) == 1 {
			url = http.BuildURL(url, "/value")
		} else {
			url = http.BuildURL(url, "/", args[1])
		}

		fmt.Println(http.Get(url, ""))
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}
