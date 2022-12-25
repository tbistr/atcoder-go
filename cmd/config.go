package cmd

import (
	"github.com/spf13/cobra"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Show global configs",
	Long:  `Show global configs.`,
	Run: func(cmd *cobra.Command, args []string) {
		exit1withE(h.ShowGlobalConfig())
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
}
