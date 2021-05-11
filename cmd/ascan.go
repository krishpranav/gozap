package cmd

import (
	"fmt"
	zap "github.com/krishpranav/gozap/pkg/zap"
	"github.com/spf13/cobra"
)

var ascanCmd = &cobra.Command{
	Use:   "ascan",
	Short: "Add ActiveScan ZAP",
	Run: func(cmd *cobra.Command, args []string) {
		if URLs != "" {
			zap.ActiveScan(URLs, apiHosts, options)
		} else {
			fmt.Println("Please input --urls flag")
		}
	},
}

func init() {
	rootCmd.AddCommand(ascanCmd)
}