package cmd

import (
	flow_module "github.com/HRKings/gitgud-cli/modules/flow"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"os"
)

var (
	branchName string
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Creates a branch with the prefixed branch type",
	Run: func(cmd *cobra.Command, args []string) {
		err := flow_module.ExecStartCommand(branchName)
		if err != nil {
			color.Red(err.Error())
			os.Exit(1)
		}
	},
}

func init() {
	flowCmd.AddCommand(startCmd)

	startCmd.Flags().StringVarP(&branchName, "branchName", "b", "", "The `<name>` of the created branch")
}
