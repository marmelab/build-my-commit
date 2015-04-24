package statehandlers

import (
	"fmt"
	"github.com/marmelab/buildmycommit/states"
	"github.com/marmelab/buildmycommit/docker"
)

// MakeAbsolutePath is the signature of a function which returns the absolute path of a given path
type MakeAbsolutePath func(path string) (string, error)

// RunDocker is the StateHandler in charge of the RunDocker state
type RunDocker struct {
	executeDocker    docker.ExecuteDocker
	makeAbsolutePath MakeAbsolutePath
}

// Handle the RunDocker state
func (stateHandler RunDocker) Handle(state states.State) (int, states.State) {
	repositoryFullPath, err := stateHandler.makeAbsolutePath(state.RepositoryPath)

	if err != nil {
		return states.InternalServerError, state
	}

	// docker run --name=test-repository-for-build-my-commit --volume=/home/gildas/projects/go/src/github.com/marmelab/build-my-commit/src/test-repository-for-build-my-commit/:/srv/ test-repository-for-build-my-commit make --file=/srv/Makefile
	err = stateHandler.executeDocker(
		"run",
		fmt.Sprintf("--name=%v", state.PushEvent.Repository.Name), // Uses the tag name we specified on build
		fmt.Sprintf("--volume=\"%v:/srv/\"", repositoryFullPath),  // Mount the repository inside the container
		state.PushEvent.Repository.Name,
		"make",
		"build")

	if err != nil {
		return states.InternalServerError, state
	}

	return states.CompareBuild, state
}

// NewRunDocker is a constructor like factory which returns a RunDocker
func NewRunDocker(executeDocker docker.ExecuteDocker, makeAbsolutePath MakeAbsolutePath) *RunDocker {
	handler := new(RunDocker)
	handler.executeDocker = executeDocker
	handler.makeAbsolutePath = makeAbsolutePath

	return handler
}
