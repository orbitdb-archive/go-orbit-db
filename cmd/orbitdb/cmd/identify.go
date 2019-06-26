package cmd

import (
	"github.com/spf13/cobra"
	"github.com/orbitdb/go-orbit-db/http"
	"fmt"
)

var identifyCmd = &cobra.Command{
	Use:   "identify",
	Short: "OrbitDB's identity settings.",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(http.Get(http.BuildURL(httpHost, "/identity"), ""))
	},
}

func init() {
	rootCmd.AddCommand(identifyCmd)
}
