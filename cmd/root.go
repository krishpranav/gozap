package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"

	version "github.com/krishpranav/gozap/pkg/version"
	zap "github.com/krishpranav/gozap/pkg/zap"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile, URLs, apiHosts, APIKey string
var options zap.OptionsZAP

var rootCmd = &cobra.Command{
	Use: "gozap",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	version.Banner()
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.mzap.yaml)")

	rootCmd.PersistentFlags().StringVar(&APIKey, "apikey", "", "ZAP API Key / if you disable apikey, not use this option")
	rootCmd.PersistentFlags().StringVar(&URLs, "urls", "", "URL list file / e.g --urls hosts.txt")
	rootCmd.PersistentFlags().StringVar(&apiHosts, "apis", "http://localhost:8090", "ZAP API Host(s) address\ne.g --apis http://localhost:8090,http://192.168.0.4:8090")

	options = zap.OptionsZAP{
		APIKey: APIKey,
		URLs:   URLs,
	}
	_ = options
}

func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".mzap" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".mzap")
	}

	viper.AutomaticEnv() 

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}