package cmd

import (
  "fmt"
  "os"
  "github.com/spf13/cobra"

  homedir "github.com/mitchellh/go-homedir"
  "github.com/spf13/viper"
)


var cfgFile string

var httpHost = "https://localhost:3000"


// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
  Use:   "orbitdb",
  Short: "OrbitDB is a decentralized database system.",
  Long: ``,
}

func Execute() {
  if err := rootCmd.Execute(); err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
}

func init() {
  cobra.OnInitialize(initConfig)

  rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.orbitdb.yaml)")
  rootCmd.PersistentFlags().StringVar(&httpHost, "http-host", httpHost, "HTTP API host address.")

  rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func initConfig() {
  if cfgFile != "" {
    viper.SetConfigFile(cfgFile)
  } else {
    home, err := homedir.Dir()
    if err != nil {
      fmt.Println(err)
      os.Exit(1)
    }

    viper.AddConfigPath(home)
    viper.SetConfigName(".orbitdb")
  }

  viper.AutomaticEnv()

  // If a config file is found, read it in.
  if err := viper.ReadInConfig(); err == nil {
    fmt.Println("Using config file:", viper.ConfigFileUsed())
  }
}
