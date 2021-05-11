package cmd

import (
	"fmt"
	zap "github.com/krishpranav/gozap/pkg/zap"
	"github.com/spf13/cobra"
)


var ajaxspiderCmd = &cobra.Command{
	Use:   "ajaxspider",
	Short: "Add AjaxSpider ZAP",
	Run: func(cmd *cobra.Command, args []string) {
		if URLs != "" {
			zap.AjaxSpider(URLs, apiHosts, options)
		} else {
			fmt.Println("Please input --urls flag")
		}
	},
}

func init() {
	rootCmd.AddCommand(ajaxspiderCmd)
}