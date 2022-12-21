package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	templateFile string
)

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new",
	Short: "create directory and templates for contest",
	Long:  `create directory and templates for contest.`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := h.NewContest(args[1], templateFile); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(newCmd)
	rootCmd.Flags().StringVar(&templateFile, "template", ".template", "template file for each problem")
}
