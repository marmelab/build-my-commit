package docker

import (
	"os"
	"os/exec"
)

var execCommand = exec.Command

// ExecuteDocker is the signature of a function which execute a docker command
type ExecuteDocker func(command string, arguments ...string) error

// Docker is a wrapper around exec to run docker commands
// Its redirects the command output to the os Std
var Docker = func(command string, arguments ...string) error {
	// Build the command
	args := []string{command}
	args = append(args, arguments...)
	cmd := execCommand("docker", args...)

	// TODO: capture output so that it may be saved later in order to report it to the user
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
