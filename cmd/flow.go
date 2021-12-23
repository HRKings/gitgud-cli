package cmd

import (
	"github.com/spf13/cobra"
)

// flowCmd represents the flow command
var flowCmd = &cobra.Command{
	Use:   "flow",
	Short: "The GitGud Flow helper",
}

func init() {
	RootCmd.AddCommand(flowCmd)
}
