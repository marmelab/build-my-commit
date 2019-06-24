package docker

import (
	"log"
	"os"
	"os/exec"
	"testing"
)

func TestCanGitCallExecCommandWithGitCommand(t *testing.T) {
	if os.Getenv("CI") == "true" {
		log.Println("Skip docker test on Travis as it does not support running docker")
		return
	}

	// Override the private package variable execCommand for testing purposes
	execCommand := func(name string, arg ...string) *exec.Cmd {
		if name != "docker" {
			t.Errorf("Expected %q to be docker", name)
		}

		return &exec.Cmd{}
	}

	docker := GetDockerCmd(execCommand)

	// Test should pass even if docker is not installed on environment so we don't store the error returned by Docker
	docker.Exec("subcommand")
}

func TestCanGitCallExecCommandWithGitCommandOptions(t *testing.T) {
	if os.Getenv("CI") == "true" {
		log.Println("Skip docker test on Travis as it does not support running docker")
		return
	}

	// Override the private package variable execCommand for testing purposes
	execCommand := func(name string, arg ...string) *exec.Cmd {
		argsLength := len(arg)

		if argsLength != 2 {
			t.Errorf("Expected arg to contains 2 options, got %v", argsLength)
		}

		if arg[0] != "subcommand" {
			t.Errorf("Expected arg[0] to equal 'subcommand', got %v", arg[0])
		}

		if arg[1] != "--option" {
			t.Errorf("Expected arg[0] to equal '--option', got %v", arg[1])
		}

		return &exec.Cmd{}
	}

	docker := GetDockerCmd(execCommand)

	// Test should pass even if docker is not installed on environment so we don't store the error returned by Docker
	docker.Exec("subcommand", "--option")
}

func TestGetDockerCmdInitializeExecCommandWithDefaults(t *testing.T) {
	docker := GetDockerCmd()

	if docker.execCommand == nil {
		t.Errorf("Expected docker.execCommand to be initialized")
	}
}
