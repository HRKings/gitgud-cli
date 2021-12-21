package utils

import (
	"bytes"
	"errors"
	"github.com/fatih/color"
	"os/exec"
	"strings"
)

func ExecGit(subcommand string, arguments ...string) (string, string, error) {
	args := append([]string{subcommand}, arguments...)
	commands := exec.Command("git", args...)

	var standardOutput bytes.Buffer
	var standardError bytes.Buffer

	commands.Stdout = &standardOutput
	commands.Stderr = &standardError

	err := commands.Run()

	return string(standardOutput.Bytes()), string(standardError.Bytes()), err
}

func CanCommit() error {
	stdout, stderr, err := ExecGit("status")
	if err != nil {
		return errors.New(stderr)
	}

	if strings.Contains(stdout, "nothing to commit, working tree clean") {
		return errors.New(stdout)
	}

	if strings.Contains(stdout, "no changes added to commit") {
		return errors.New(stdout)
	}

	return nil
}

func HandleGitError(stdout string, stderr string, err error) error {
	// Print the stdout and stderr in case of an error
	if err != nil && stderr != "" {
		color.Red(stderr)
		return err
	} else if err != nil && stdout != "" {
		color.Red(stdout)
		return err
	} else if err != nil {
		return err
	}

	return nil
}
