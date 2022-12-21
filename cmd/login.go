package cmd

import (
	"github.com/spf13/cobra"
)

var (
	sessionFile string
)

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "login to AtCoder",
	Long:  `Login to AtCoder.`,
	Run: func(cmd *cobra.Command, args []string) {
		h.Login(sessionFile)
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)

	loginCmd.Flags().StringVarP(&sessionFile, "session", "s", ".atcoder_session", "Filename for keep session.")
}
