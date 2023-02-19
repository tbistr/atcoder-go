package cmd

import (
	"github.com/spf13/cobra"
)

// testCmd represents the test command
var testCmd = &cobra.Command{
	Use:   "test",
	Short: "test program",
	Long:  `test program`,
	Run: func(cmd *cobra.Command, args []string) {
		exit1withE(h.Test())
	},
}

func init() {
	rootCmd.AddCommand(testCmd)
}
