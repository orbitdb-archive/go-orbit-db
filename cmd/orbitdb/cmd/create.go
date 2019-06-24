package cmd

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
	"github.com/orbitdb/go-orbit-db/api"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Creates a database or opens it if it already exists.",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 2 {
			fmt.Println("Usage: orbitdb create name type")
			os.Exit(1)
		}

		params := make(map[string]string)
		params["create"] = "true"
		params["type"] = args[1]

		paramsString := api.MapToJSONString(params)

		fmt.Println(api.BuildURL(apiHost, "/db/", args[0]))
		fmt.Println(api.Post(api.BuildURL(apiHost, "/db/", args[0]), paramsString))
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
}
