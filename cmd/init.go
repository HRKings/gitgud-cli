package cmd

import (
	flow_module "github.com/HRKings/gitgud-cli/modules/flow"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"os"
)

var (
	remote        string
	defaultBranch string
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initializes a local git repository with a initial commit and stable branch",
	Long: `If a remote is provided, it will automatically push the repo to the remote`,
	Run: func(cmd *cobra.Command, args []string) {
		err := flow_module.ExecuteInitCommand(remote, defaultBranch)
		if err != nil {
			color.Red(err.Error())
			os.Exit(1)
		}
	},
}

func init() {
	flowCmd.AddCommand(initCmd)

	initCmd.Flags().StringVarP(&remote, "remote", "r", "", "The `<url>` of the remote repository")
	initCmd.Flags().StringVar(&remote, "defaultBranch", "master", "The `<name>` of the default branch")
}
