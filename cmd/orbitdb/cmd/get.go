package cmd

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
	"github.com/orbitdb/go-orbit-db/http"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Gets a record from the database.",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		//var data = ""

		if len(args) < 1 {
			fmt.Println("Usage: orbitdb get name")
			os.Exit(1)
		}

		/*if len(args) == 2 {
			data = args[1]
		}*/

		fmt.Println(http.Get(http.BuildURL(httpHost, "/db/", args[0], "/get")))
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}
