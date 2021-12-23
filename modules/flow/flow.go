package flow_module

import (
	"fmt"
	"github.com/HRKings/gitgud-cli/utils"
	"github.com/fatih/color"
)

func ExecuteInitCommand(remote string, defaultBranch string) error {
	// If we don't have a local repository, create one
	err := utils.HasRepository()
	if err != nil {
		stdout, stderr, gitErr := utils.ExecGit("init")
		err := utils.HandleGitError(stdout, stderr, gitErr)
		if err != nil {
			return err
		}

		color.Blue("> Created a new Git repository...")
	}

	// Create the first commit
	stdout, stderr, gitErr := utils.ExecGit("commit",
		"--allow-empty", "-m", "[misc] Initial commit")
	err = utils.HandleGitError(stdout, stderr, gitErr)
	if err != nil {
		return err
	}
	color.Blue("> Created the initial commit...")

	// Crete the stable branch
	stdout, stderr, gitErr = utils.ExecGit("checkout", "-b", "stable")
	err = utils.HandleGitError(stdout, stderr, gitErr)
	if err != nil {
		return err
	}
	color.Blue("> Created branch stable...")

	// If we have a remote
	if remote != "" {
		// Add the remote to the local repo
		stdout, stderr, gitErr = utils.ExecGit("remote", "add", "origin", remote)
		err = utils.HandleGitError(stdout, stderr, gitErr)
		if err != nil {
			return err
		}
		color.Blue("> Added remote 'origin' pointing to '%s'", remote)

		// Push the default branch to the remote
		stdout, stderr, gitErr = utils.ExecGit("push", "-u", "origin", defaultBranch)
		err = utils.HandleGitError(stdout, stderr, gitErr)
		if err != nil {
			return err
		}
		color.Blue("> Pushed '%s' to 'origin'", defaultBranch)

		// Push the stable branch to the remote
		stdout, stderr, gitErr = utils.ExecGit("push", "-u", "origin", "stable")
		err = utils.HandleGitError(stdout, stderr, gitErr)
		if err != nil {
			return err
		}
		color.Blue("> Pushed 'stable' to 'origin'")
	}

	// Go back to the default branch
	stdout, stderr, gitErr = utils.ExecGit("checkout", defaultBranch)
	err = utils.HandleGitError(stdout, stderr, gitErr)
	if err != nil {
		return err
	}
	color.Blue("> Changed to default branch '%s'", defaultBranch)

	// Print the success message
	color.Green("Fully initialized the repository")
	return nil
}

func ExecStartCommand(branchName string) error {
	err := utils.HasRepository()
	if err != nil {
		return err
	}

	if branchName == "" {
		branchName, err = EnterBranchName()
		if err != nil {
			return err
		}
	}

	branchType, err := EnterBranchType()
	if err != nil {
		return err
	}

	branch := fmt.Sprintf("%s%s", branchType, branchName)
	stdout, stderr, gitErr := utils.ExecGit("checkout", "-b", branch)
	err = utils.HandleGitError(stdout, stderr, gitErr)
	if err != nil {
		return err
	}

	// Print the stdout of git
	color.Green("Created and switched to branch '%s'", branch)
	return nil
}