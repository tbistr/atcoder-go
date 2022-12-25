package cmd

import (
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
		// TODO: 引数のサイズチェック
		exit1withE(h.NewContest(args[0], templateFile))
	},
}

func init() {
	rootCmd.AddCommand(newCmd)

	newCmd.Flags().StringVar(&templateFile, "template", ".template", "template file for each problem")
}
