package git

import (
	"fmt"
	"os/exec"
	"testing"
)

func TestGitCallExecCommandWithGitCommand(t *testing.T) {
	// Override the private package variable execCommand for testing purposes
	execCommand := func(name string, arg ...string) *exec.Cmd {
		if name != "git" {
			t.Errorf("Expected %v to be git", name)
		}

		return &exec.Cmd{}
	}

	git := GetGitCmd(execCommand)

	// Test should pass even if git is not installed on environment so we don't store the error returned by Git
	git.Exec("version", "")
}

func TestGitCallExecCommandWithGitCommandOptions(t *testing.T) {
	// Override the private package variable execCommand for testing purposes
	execCommand := func(name string, arg ...string) *exec.Cmd {
		argsLength := len(arg)

		if argsLength != 2 {
			t.Errorf("Expected arg to contains 2 options, got %v", argsLength)
		}

		if arg[0] != "version" {
			t.Errorf("Expected arg[0] to equal 'version', got %v", arg[0])
		}

		if arg[1] != "--option" {
			t.Errorf("Expected arg[0] to equal '--option', got %v", arg[1])
		}

		return &exec.Cmd{}
	}

	git := GetGitCmd(execCommand)

	// Test should pass even if git is not installed on environment so we don't store the error returned by Git
	git.Exec("version", "--option")
}

func TestGitWithContextPassesRepositoryOptionsBeforeGitCommand(t *testing.T) {
	repositoryPath := "repository-path"
	gitDir := fmt.Sprintf("--git-dir=%v/.git", repositoryPath)
	workTree := fmt.Sprintf("--work-tree=%v", repositoryPath)

	// Override the private package variable execCommand for testing purposes
	execCommand := func(name string, arg ...string) *exec.Cmd {
		if name != "git" {
			t.Errorf("Expected %v to be git", name)
		}

		argsLength := len(arg)

		if argsLength != 3 {
			t.Errorf("Expected arg to contains 3 element, got %v", argsLength)
		}

		if arg[0] != gitDir {
			t.Errorf("Expected arg[0] to equal %v, got %v", gitDir, arg[0])
		}

		if arg[1] != workTree {
			t.Errorf("Expected arg[1] to equal %v, got %v", workTree, arg[1])
		}

		return &exec.Cmd{}
	}

	git := GetGitCmd(execCommand)

	// Test should pass even if git is not installed on environment so we don't store the error returned by Git
	git.ExecInContext("version", repositoryPath)
}
