package main

import (
	"errors"
	"os"
	"os/exec"
	"regexp"
)

var gitCloneUrlRegexp = regexp.MustCompile(`^https:\/\/github\.com\/.+\/(.+)\.git$`)

func CloneRepository(gitCloneUrl string) (path string, err error) {
	// Validate the git url
	matches := gitCloneUrlRegexp.FindStringSubmatch(gitCloneUrl)

	if gitCloneUrl == "" {
		return "", errors.New("Invalid git clone url")
	}

	// Build the command
	args := []string{"clone", gitCloneUrl}
	cmd := exec.Command("git", args...)

	// TODO: capture output so that it may be saved later in order to report it to the user
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()

	if err != nil {
		return "", err
	}

	outputPath := matches[1]
	return outputPath, err
}
