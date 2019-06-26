package cmd

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
	"github.com/orbitdb/go-orbit-db/http"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Gets a record from the database matching id.",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 2 {
			fmt.Println("Usage: orbitdb get name id")
			os.Exit(1)
		}

		fmt.Println(http.Get(http.BuildURL(httpHost, "/db/", args[0], "/", args[1]), ""))
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}
