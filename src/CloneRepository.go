package main

import (
	"errors"
	"os"
	"os/exec"
	"regexp"
)

var gitCloneUrlRegexp = regexp.MustCompile(`^https:\/\/github\.com\/.+\/(.+)\.git$`)

const GIT_COMMAND = "/usr/bin/git"
const GIT_CLONE_COMMAND = "clone"

func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func CloneRepository(gitCloneUrl string) (path string, err error) {
	// Validate the git url
	matches := gitCloneUrlRegexp.FindStringSubmatch(gitCloneUrl)

	if gitCloneUrl == "" || len(matches) < 2 {
		return "", errors.New("Invalid git clone url")
	}

	outputPath := matches[1]

	// Build the command
	args := []string{GIT_CLONE_COMMAND, gitCloneUrl}
	cmd := exec.Command(GIT_COMMAND, args...)

	err = cmd.Start()

	if err != nil {
		return "", err
	}

	err = cmd.Wait()

	if err != nil {
		return "", err
	}

	directoryExist, err := exists(outputPath)
	if !directoryExist {
		return "", err
	}

	return outputPath, err
}
