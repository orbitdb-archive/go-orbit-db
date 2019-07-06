package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/orbitdb/go-orbit-db/http"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all registered databases.",
	Long: `
The list command lists all registered databases. A registered database is one
that has been created and/or loaded using the create command.`,
	Run: func(cmd *cobra.Command, args []string) {
		var r http.Request

		r.SetURL(httpHost, "/dbs")

		fmt.Println(r.Get())
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
