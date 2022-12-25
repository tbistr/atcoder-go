package cmd

import (
	"github.com/spf13/cobra"
)

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "login to AtCoder",
	Long:  `Login to AtCoder.`,
	Run:   func(cmd *cobra.Command, args []string) { exit1withE(h.Login()) },
}

func init() {
	rootCmd.AddCommand(loginCmd)
}
