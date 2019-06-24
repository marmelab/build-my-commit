package statehandlers

import (
	"github.com/marmelab/buildmycommit/states"
)

// EndRequest is the StateHandler in charge of the EndRequest state
type EndRequest struct {
}

// Handle the EndRequest state
func (stateHandler EndRequest) Handle(state states.State) (int, states.State) {
	state.ResponseWriter.WriteHeader(state.Status)
	state.ResponseWriter.Write([]byte(state.StatusMessage))

	if state.ShouldCleanGit {
		return states.CleanRepository, state
	}

	if state.ShouldCleanDocker {
		return states.CleanDocker, state
	}

	return states.RequestHandled, state
}

// NewEndRequest is a constructor like factory which returns a EndRequest
func NewEndRequest() *EndRequest {
	handler := new(EndRequest)

	return handler
}
