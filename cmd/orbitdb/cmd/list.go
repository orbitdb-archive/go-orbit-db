package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/orbitdb/go-orbit-db/api"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all registered databases.",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(api.Get(api.BuildURL(apiHost, "/dbs")))
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
