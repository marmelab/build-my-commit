package main

import (
	"os"
	"os/exec"
)

func docker(command string, arguments ...string) error {
	// Build the command
	args := []string{command}
	args = append(args, arguments...)
	cmd := exec.Command("docker", args...)

	// TODO: capture output so that it may be saved later in order to report it to the user
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
