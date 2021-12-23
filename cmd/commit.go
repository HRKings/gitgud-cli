package cmd

import (
	"github.com/HRKings/gitgud-cli/modules/commit"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"os"
)

var (
	message string
	amend   bool
	all     bool
	quick   bool
	scope   string
)

// commitCmd represents the commit command
var commitCmd = &cobra.Command{
	Use:     "commit",
	Aliases: []string{"c"},
	Short:   "Generates a commit with a message that that follows the GitGud spec",
	Run: func(cmd *cobra.Command, args []string) {
		err := commit_module.ExecCommand(all, amend, message, scope, quick, "", "", "", false)
		if err != nil {
			color.Red(err.Error())
			os.Exit(1)
		}
	},
}

func init() {
	RootCmd.AddCommand(commitCmd)

	commitCmd.PersistentFlags().BoolVarP(&all, "all", "a", false, "Tell the command to automatically stage files that have been modified and deleted, but new files you have not told Git about are not affected.")
	commitCmd.PersistentFlags().BoolVar(&amend, "amend", false, "Replace the tip of the current branch by creating a new commit.")
	commitCmd.PersistentFlags().BoolVarP(&quick, "quick", "q", false, "Don't ask for missing parts")

	commitCmd.PersistentFlags().StringVarP(&message, "message", "m", "", "Use the given `<msg>` as the commit Subject.")
	commitCmd.PersistentFlags().StringVarP(&scope, "scope", "s", "", "`<scope>` is the area of the code that the commit affects.")
}
