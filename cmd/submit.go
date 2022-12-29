package cmd

import (
	"github.com/spf13/cobra"
)

// submitCmd represents the submit command
var submitCmd = &cobra.Command{
	Use:   "submit",
	Short: "submit program",
	Long:  `submit program`,
	Run: func(cmd *cobra.Command, args []string) {
		exit1withE(h.Submit())
	},
}

func init() {
	rootCmd.AddCommand(submitCmd)
}
