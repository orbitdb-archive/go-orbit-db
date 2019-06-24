package cmd

import (
	"github.com/spf13/cobra"
	"github.com/orbitdb/go-orbit-db/api"
	"fmt"
)

var identityCmd = &cobra.Command{
	Use:   "identity",
	Short: "OrbitDB's identity settings.",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(api.Get(api.BuildURL(apiHost, "/identity")))
	},
}

func init() {
	rootCmd.AddCommand(identityCmd)
}
