package cmd

import (
	"fmt"

	zap "github.com/krishpranav/gozap/pkg/zap"
	"github.com/spf13/cobra"
)

// spiderCmd represents the spider command
var spiderCmd = &cobra.Command{
	Use:   "spider",
	Short: "Add ZAP spider",
	Run: func(cmd *cobra.Command, args []string) {
		if URLs != "" {
			zap.Spider(URLs, apiHosts, options)
		} else {
			fmt.Println("Please input --urls flag")
		}
	},
}

func init() {
	rootCmd.AddCommand(spiderCmd)
}
