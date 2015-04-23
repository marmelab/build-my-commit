package main

import (
	"fmt"
	"path/filepath"
)

var filepathAbs = filepath.Abs

func HandleStateRunDocker(state State) (int, State, error) {
	repositoryFullPath, err := filepathAbs(state.RepositoryPath)

	if err != nil {
		return StateEndRequest, state, err
	}

	// docker run --name=test-repository-for-build-my-commit --volume=/home/gildas/projects/go/src/github.com/marmelab/build-my-commit/src/test-repository-for-build-my-commit/:/srv/ test-repository-for-build-my-commit make --file=/srv/Makefile
	err = docker(
		"run",
		fmt.Sprintf("--name=%v", state.PushEvent.Repository.Name), // Uses the tag name we specified on build
		fmt.Sprintf("--volume=\"%v:/srv/\"", repositoryFullPath),  // Mount the repository inside the container
		state.PushEvent.Repository.Name,
		"make",
		"build")

	if err != nil {
		return StateEndRequest, state, err
	}

	return StateCompareOutput, state, nil
}
