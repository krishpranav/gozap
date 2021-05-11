package cmd

import (
	"fmt"

	version "github.com/krishpranav/gozap/pkg/version"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(version.VERSION)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
