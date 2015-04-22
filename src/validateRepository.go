package main

import (
	"path"
	"strings"
)

func validateRepository(repositoryPath string) (bool, error) {
	dockerFileFullPath := path.Join(repositoryPath, dockerFilePath)

	dockerFileExists, err := exists(dockerFileFullPath)

	if err != nil {
		return false, err
	}

	if dockerFileExists {
		// Get the last commit message without pretty formatting
		output, err := gitWithContext(
			"log",
			repositoryPath,
			"-1",
			"--pretty=%B")

		if err != nil {
			return false, err
		}

		// If the last commit message isn't our standard commit message
		if !strings.Contains(output, commitMessage) {
			return true, nil
		}
	}

	return false, nil
}
