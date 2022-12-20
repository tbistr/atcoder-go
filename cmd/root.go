package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/tbistr/atcoder-go/atcodergo"
)

var client *atcodergo.Client

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "atgo",
	Short: "",
	Long:  ``,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initClient)
}

func initClient() {
	var err error
	client, err = atcodergo.NewClient()
	if err != nil {
		fmt.Fprintln(os.Stderr, "failed to init atcoder-go library")
		os.Exit(1)
	}
}
