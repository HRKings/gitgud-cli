package cmd

import (
	commit_module "github.com/HRKings/gitgud-cli/modules/commit"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"os"
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:     "generate",
	Aliases: []string{"gen"},
	Short:   "Generates a commit message compliant with the GitGud spec",
	Run: func(cmd *cobra.Command, args []string) {
		message, err := commit_module.BuildFullCommitMessage(message, scope, quick, body, closes, see)
		if err != nil {
			color.Red(err.Error())
			os.Exit(1)
		}

		color.Blue(message)
	},
}

func init() {
	commitCmd.AddCommand(generateCmd)

	generateCmd.Flags().StringVarP(&body, "body", "b", "", "The `<body>` of this commit.")
	generateCmd.Flags().StringVarP(&closes, "closes", "c", "", "`<#issue1, #issueN>` is a comma separated list of issues that this commit closes.")
	generateCmd.Flags().StringVarP(&see, "see", "r", "", "`<#issue1, #issueN>` is a comma separated list of issues that this commit references.")
}
