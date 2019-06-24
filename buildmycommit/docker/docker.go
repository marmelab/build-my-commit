package docker

import "os"
import "os/exec"

//var execCommand = exec.Command

// ExecuteDocker is the signature of a function which execute a docker command
type ExecuteDocker func(command string, arguments ...string) error

// ExecuteCommand is the signature of a function which execute a docker command
type ExecuteCommand func(command string, arguments ...string) *exec.Cmd

// Docker is a wrapper around os.exec.Command to execute docker command
type Docker struct {
	execCommand ExecuteCommand
}

// Exec execute the specified docker command
func (d Docker) Exec(command string, arguments ...string) error {
	// Build the command
	args := []string{command}
	args = append(args, arguments...)
	cmd := d.execCommand("docker", args...)

	// TODO: capture output so that it may be saved later in order to report it to the user
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

// GetDockerCmd returns an object which can execute docker command
func GetDockerCmd(execCommand ...ExecuteCommand) Docker {
	if len(execCommand) == 0 {
		execCommand = make([]ExecuteCommand, 1)
		execCommand[0] = exec.Command
	}

	return Docker{execCommand: execCommand[0]}
}
