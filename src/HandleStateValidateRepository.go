package main

import (
	"net/http"
	"path"
	"strings"
)

var pathJoin = path.Join

const StatusMessageForStateValidateRepositoryCommitDetected = "Discarded: detected commit by automatic build"
const StatusMessageForStateValidateRepositoryNoDockerfile = "Discarded: no build.Dockerfile for automatic build"

func HandleStateValidateRepository(state State) (int, State, error) {
	dockerFileFullPath := pathJoin(state.RepositoryPath, DockerFilePath)

	dockerFileExists, err := exists(dockerFileFullPath)

	if err != nil {
		return StateEndRequest, state, err
	}

	if dockerFileExists {
		// Get the last commit message without pretty formatting
		output, err := gitWithContext(
			"log",
			state.RepositoryPath,
			"-1",
			"--pretty=%B")

		if err != nil {
			return StateEndRequest, state, err
		}

		// If the last commit message is our standard commit message, we end here
		if strings.Contains(output, CommitMessage) {
			state.Status = http.StatusOK
			state.StatusMessage = StatusMessageForStateValidateRepositoryCommitDetected
			return StateEndRequest, state, nil
		}

		return StateBuildDocker, state, nil
	}

	state.Status = http.StatusOK
	state.StatusMessage = StatusMessageForStateValidateRepositoryNoDockerfile
	return StateEndRequest, state, nil
}
