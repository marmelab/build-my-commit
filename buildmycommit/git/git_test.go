package git

import (
	"fmt"
	"os/exec"
	"testing"
)

func TestGitCallExecCommandWithGitCommand(t *testing.T) {
	// Override the private package variable execCommand for testing purposes
	execCommand = func(name string, arg ...string) *exec.Cmd {
		if name != "git" {
			t.Errorf("Expected %v to be git", name)
		}

		return &exec.Cmd{}
	}

	// Dirty hack to prevent compilation warning about execCommand not being used
	_ = execCommand

	// Test should pass even if git is not installed on environment so we don't store the error returned by Git
	Git("version", "")
}

/*
type fakeCmd struct{}

func (cmd *fakeCmd) Output() ([]byte, error) {
	return []byte("42"), nil
}

func TestGitReturnsExecCommandOutput(t *testing.T) {
	expectedOutput := "42"

	// Override the private package variable execCommand for testing purposes
	execCommand = func(name string, arg ...string) *fakeCmd {
		if name != "git" {
			t.Errorf("Expected %v to be git", name)
		}

		cmd := fakeCmd{}

		return &cmd
	}

	// Dirty hack to prevent compilation warning about execCommand not being used
	_ = execCommand

	// Test should pass even if git is not installed on environment so we don't store the error returned by Git
	output, _ := Git("version", "")

	if output != expectedOutput {
		t.Errorf("Expected output to equal '%v', got '%v'", expectedOutput, output)
	}
}
*/

func TestGitCallExecCommandWithGitCommandOptions(t *testing.T) {
	// Override the private package variable execCommand for testing purposes
	execCommand = func(name string, arg ...string) *exec.Cmd {
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

	// Dirty hack to prevent compilation warning about execCommand not being used
	_ = execCommand

	// Test should pass even if git is not installed on environment so we don't store the error returned by Git
	Git("version", "--option")
}

func TestGitWithContextPassesRepositoryOptionsBeforeGitCommand(t *testing.T) {
	repositoryPath := "repository-path"
	gitDir := fmt.Sprintf("--git-dir=%v/.git", repositoryPath)
	workTree := fmt.Sprintf("--work-tree=%v", repositoryPath)

	// Override the private package variable execCommand for testing purposes
	execCommand = func(name string, arg ...string) *exec.Cmd {
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

	// Dirty hack to prevent compilation warning about execCommand not being used
	_ = execCommand

	// Test should pass even if git is not installed on environment so we don't store the error returned by Git
	GitWithContext("version", repositoryPath)
}
