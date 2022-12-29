// cmd is package to parse args and options.
// Does not have any implementation.
package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/tbistr/atcoder-go/atcodergo"
	"github.com/tbistr/atcoder-go/cmd/handler"
)

var (
	h          *handler.Handler
	configFile string
)

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

	rootCmd.PersistentFlags().StringVar(
		&configFile, "config",
		configDir("config.json"),
		"Global config file(json).")
}

func initClient() {
	var err error
	h, err = handler.New(
		configFile,
		&handler.GlobalConfig{
			SessionFile:          configDir(".atcoder_session"),
			TemplateCmdName:      "cat",
			TemplateCmdArgs:      []string{},
			TemplateCmdJsonInput: false,
			TemplateFile:         configDir("template"),
			MainFileName:         "main.go",
			DefaultLanguage: atcodergo.Language{
				Value:    "4026",
				Datamime: "text/x-go",
				Text:     "Go (1.14.1)"},
		},
	)
	exit1withE(err)
}
