package statehandlers

import (
	"github.com/marmelab/buildmycommit/states"
	"github.com/marmelab/buildmycommit/docker"
)

// CleanDocker is the StateHandler in charge of the CleanDocker state
type CleanDocker struct {
	executeDocker docker.ExecuteDocker
}

// Handle the CleanDocker state
func (stateHandler CleanDocker) Handle(state states.State) (int, states.State) {
	err := stateHandler.executeDocker("rm", state.PushEvent.Repository.Name)

	if err != nil {
		return states.InternalServerError, state
	}

	state.ShouldCleanDocker = false
	return states.EndRequest, state
}

// NewCleanDocker is a constructor like factory which returns a CleanDocker
func NewCleanDocker(executeDocker docker.ExecuteDocker) *CleanDocker {
	handler := new(CleanDocker)
	handler.executeDocker = executeDocker

	return handler
}
