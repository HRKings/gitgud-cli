package cmd

import (
	"github.com/HRKings/gitgud-cli/modules/commit"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"os"
)

var (
	body   string
	closes string
	see    string
)

// fullCmd represents the full command
var fullCmd = &cobra.Command{
	Use:     "full",
	Aliases: []string{"f"},
	Short:   "Generates a commit with a full message that follows the GitGud spec",
	Run: func(cmd *cobra.Command, args []string) {
		err := commit_module.ExecCommand(all, amend, message, scope, quick, body, closes, see, true)
		if err != nil {
			color.Red(err.Error())
			os.Exit(1)
		}
	},
}

func init() {
	commitCmd.AddCommand(fullCmd)

	fullCmd.Flags().StringVarP(&body, "body", "b", "", "The `<body>` of this commit.")
	fullCmd.Flags().StringVarP(&closes, "closes", "c", "", "`<#issue1, #issueN>` is a comma separated list of issues that this commit closes.")
	fullCmd.Flags().StringVarP(&see, "see", "r", "", "`<#issue1, #issueN>` is a comma separated list of issues that this commit references.")
}
