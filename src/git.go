package main

import (
	"fmt"
	"os"
	"os/exec"
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

	return string(output), nil
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

	return string(output), nil
}

var (
	git            = _git
	gitWithContext = _gitWithContext
)
