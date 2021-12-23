package commit_module

import (
	"github.com/HRKings/gitgud-cli/utils"
	"github.com/fatih/color"
)

func ExecCommand(all bool, amend bool, message string, scope string, quick bool, body string, closes string, see string, full bool) error {
	// Verify if there is changes to commit
	err := utils.CanCommit()
	if err != nil {
		return err
	}

	// Build a full message or a simple one
	if !full {
		message, err = BuildCommitMessage(message,
			scope, quick)
	} else {
		message, err = BuildFullCommitMessage(message,
			scope, quick, body,
			closes, see)
	}
	if err != nil {
		return err
	}

	// Execute the commit is everything is all right
	stdout, stderr, gitErr := utils.DoCommit(message, all, amend)
	err = utils.HandleGitError(stdout, stderr, gitErr)
	if err != nil {
		return err
	}

	// Print the stdout of git
	color.Green(stdout)
	return nil
}
