package main

import (
	"fmt"
	"path"
)

func HandleStateBuildDocker(state State) (int, State, error) {
	state.ShouldCleanDocker = true

	err := docker(
		"build",
		fmt.Sprintf("--tag=%v", state.PushEvent.Repository.Name),                  // This allows us to find that container for run
		fmt.Sprintf("--file=%v", path.Join(state.RepositoryPath, DockerFilePath)), // Specify the file as we don't use default Dockerfile
		state.RepositoryPath)

	if err != nil {
		return StateEndRequest, state, err
	}

	return StateRunDocker, state, nil
}
