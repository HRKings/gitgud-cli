package utils

import (
	"bytes"
	"errors"
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
