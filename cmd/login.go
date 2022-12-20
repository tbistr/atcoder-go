/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"syscall"

	"github.com/spf13/cobra"
	"golang.org/x/term"
)

var (
	sessionFile string
)

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "login to AtCoder",
	Long:  `Login to AtCoder.`,
	Run:   login,
}

func init() {
	rootCmd.AddCommand(loginCmd)

	loginCmd.Flags().StringVarP(&sessionFile, "session", "s", ".atcoder_session", "Filename for keep session.")
}

func login(cmd *cobra.Command, args []string) {
	if client.LoginWithSession(sessionFile) != nil {
		// hear username, password
		var u, p string
		fmt.Fprint(os.Stderr, "enter username:")
		fmt.Scanln(&u)
		fmt.Fprint(os.Stderr, "enter password:")
		b, _ := term.ReadPassword(int(syscall.Stdin))
		p = string(b)
		fmt.Fprint(os.Stderr, "\n")

		if err := client.LoginWithNewSession(u, p, sessionFile); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}
