package main

import (
	"path"
)

func validateRepository(repositoryPath string) (shouldProcess bool, err error) {
	dockerFileFullPath := path.Join(repositoryPath, dockerFilePath)

	dockerFileExists, err := exists(dockerFileFullPath)

	if err != nil {
		return false, err
	}

	if dockerFileExists {
		return true, nil
	}

	return false, nil
}
