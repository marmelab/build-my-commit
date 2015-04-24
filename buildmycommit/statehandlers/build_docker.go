package statehandlers

import (
	"fmt"
	"github.com/marmelab/buildmycommit/states"
	"github.com/marmelab/buildmycommit/docker"
	"path"
)

// BuildDocker is the StateHandler in charge of the BuildDocker state
type BuildDocker struct {
	executeDocker  docker.ExecuteDocker
	dockerFilePath string
}

// Handle the BuildDocker state
func (stateHandler BuildDocker) Handle(state states.State) (int, states.State) {
	state.ShouldCleanDocker = true

	err := stateHandler.executeDocker(
		"build",
		// This allows us to find that container for run
		fmt.Sprintf("--tag=%v", state.PushEvent.Repository.Name),
		// Specify the file as we don't use default Dockerfile
		fmt.Sprintf("--file=%v", path.Join(state.RepositoryPath, stateHandler.dockerFilePath)),
		state.RepositoryPath)

	if err != nil {
		return states.InternalServerError, state
	}

	return states.RunDocker, state
}

// NewBuildDocker is a constructor like factory which returns a BuildDocker
func NewBuildDocker(executeDocker docker.ExecuteDocker, dockerFilePath string) *BuildDocker {
	handler := new(BuildDocker)
	handler.executeDocker = executeDocker
	handler.dockerFilePath = dockerFilePath

	return handler
}
