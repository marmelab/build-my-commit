package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

var _git = func(command string, arguments ...string) (string, error) {
	// Build the command
	args := []string{command}
	args = append(args, arguments...)
	cmd := exec.Command("git", args...)

	cmd.Stderr = os.Stderr

	output, err := cmd.Output()

	if err != nil {
		return "", err
	}

	strOutput := string(output)
	strOutput = strings.Replace(strOutput, "\n", "", -1)
	return strOutput, nil
}

var _gitWithContext = func(command string, contextPath string, arguments ...string) (string, error) {
	gitDir := fmt.Sprintf("--git-dir=%v/.git", contextPath)
	workTree := fmt.Sprintf("--work-tree=%v", contextPath)

	// Build the command
	args := []string{gitDir, workTree, command}
	args = append(args, arguments...)

	cmd := exec.Command("git", args...)

	cmd.Stderr = os.Stderr

	output, err := cmd.Output()

	if err != nil {
		return "", err
	}

	strOutput := string(output)
	strOutput = strings.Replace(strOutput, "\n", "", -1)
	return strOutput, nil
}

var _getCommitMessage = func(hash string) (string, error) {
	if hash == "" {
		return "", errors.New("hash cannot be empty")
	}

	return commitMessage + " " + hash, nil
}

var (
	git              = _git
	gitWithContext   = _gitWithContext
	getCommitMessage = _getCommitMessage
)
