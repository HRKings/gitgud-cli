package utils

import (
	"bytes"
	"os/exec"
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
