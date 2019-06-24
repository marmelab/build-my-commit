package git

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// ExecuteCommand is the signature of a function which execute a docker command
type ExecuteCommand func(command string, arguments ...string) *exec.Cmd

// ExecuteGit is the signature of a function which execute a git command
type ExecuteGit func(command string, arguments ...string) (string, error)

// ExecuteGitWithContext is the signature of a function which execute a git command in the context of a repository
type ExecuteGitWithContext func(command string, repositoryPath string, arguments ...string) (string, error)

// Git is a wrapper around exec to run git commands
type Git struct {
	execCommand ExecuteCommand
}

// Exec execute a git command
// Its redirects the command Stderr to the os Stderr and returns the command output as string (removing all new line)
func (g Git) Exec(command string, arguments ...string) (string, error) {
	// Build the command
	args := []string{command}
	args = append(args, arguments...)

	cmd := g.execCommand("git", args...)

	cmd.Stderr = os.Stderr

	output, err := cmd.Output()

	if err != nil {
		return "", err
	}

	strOutput := string(output)
	strOutput = strings.Replace(strOutput, "\n", "", -1)
	return strOutput, nil
}

// ExecInContext execute a git command in the context of a repossitory
// Its redirects the command Stderr to the os Stderr and returns the command output as string (removing all new line)
func (g Git) ExecInContext(command string, repositoryPath string, arguments ...string) (string, error) {
	// Build git options to run in the repository context
	gitDir := fmt.Sprintf("--git-dir=%v/.git", repositoryPath)
	workTree := fmt.Sprintf("--work-tree=%v", repositoryPath)

	// Build the command
	args := []string{gitDir, workTree, command}
	args = append(args, arguments...)

	cmd := g.execCommand("git", args...)

	cmd.Stderr = os.Stderr

	output, err := cmd.Output()

	if err != nil {
		return "", err
	}

	strOutput := string(output)
	strOutput = strings.Replace(strOutput, "\n", "", -1)
	return strOutput, nil
}

// GetGitCmd returns an object which can execute git commands
func GetGitCmd(execCommand ...ExecuteCommand) Git {
	if len(execCommand) == 0 {
		execCommand = make([]ExecuteCommand, 1)
		execCommand[0] = exec.Command
	}

	return Git{execCommand: execCommand[0]}
}
