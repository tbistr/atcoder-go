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
		if len(args) != 1 {
			fmt.Fprintln(os.Stderr, "Too few args. Assumes 1.")
			os.Exit(1)
		}
		exit1withE(h.NewContest(args[0], templateFile))
	},
}

func init() {
	rootCmd.AddCommand(newCmd)

	newCmd.Flags().StringVar(&templateFile, "template", ".template", "template file for each problem")
}
