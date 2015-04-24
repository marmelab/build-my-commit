package git

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

var execCommand = exec.Command

// ExecuteGit is the signature of a function which execute a git command
type ExecuteGit func(command string, arguments ...string) (string, error)

// ExecuteGitWithContext is the signature of a function which execute a git command in the context of a repository
type ExecuteGitWithContext func(command string, repositoryPath string, arguments ...string) (string, error)

// Git is a wrapper around exec to run git commands
// Its redirects the command Stderr to the os Stderr and returns the command output as string (removing all new line)
var Git = func(command string, arguments ...string) (string, error) {
	// Build the command
	args := []string{command}
	args = append(args, arguments...)

	cmd := execCommand("git", args...)

	cmd.Stderr = os.Stderr

	output, err := cmd.Output()

	if err != nil {
		return "", err
	}

	strOutput := string(output)
	strOutput = strings.Replace(strOutput, "\n", "", -1)
	return strOutput, nil
}

// GitWithContext is a wrapper around exec to run git commands in the context of a repossitory
// Its redirects the command Stderr to the os Stderr and returns the command output as string (removing all new line)
var GitWithContext = func(command string, repositoryPath string, arguments ...string) (string, error) {
	// Build git options to run in the repository context
	gitDir := fmt.Sprintf("--git-dir=%v/.git", repositoryPath)
	workTree := fmt.Sprintf("--work-tree=%v", repositoryPath)

	// Build the command
	args := []string{gitDir, workTree, command}
	args = append(args, arguments...)

	cmd := execCommand("git", args...)

	cmd.Stderr = os.Stderr

	output, err := cmd.Output()

	if err != nil {
		return "", err
	}

	strOutput := string(output)
	strOutput = strings.Replace(strOutput, "\n", "", -1)
	return strOutput, nil
}
