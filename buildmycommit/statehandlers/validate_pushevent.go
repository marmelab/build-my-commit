package statehandlers

import (
	"github.com/marmelab/buildmycommit/states"
	"net/http"
)

// ValidatePushEvent is the StateHandler in charge of the ValidatePushEvent state
type ValidatePushEvent struct {
	repositoryBranchToConsider   string
	statusMessageOnInvalidBranch string
}

// Handle the ValidatePushEvent state
func (stateHandler ValidatePushEvent) Handle(state states.State) (int, states.State) {
	if state.PushEvent.Ref == stateHandler.repositoryBranchToConsider {
		return states.CloneRepository, state
	}

	state.Status = http.StatusOK
	state.StatusMessage = stateHandler.statusMessageOnInvalidBranch
	return states.EndRequest, state
}

// NewValidatePushEvent is a constructor like factory which returns a ValidatePushEvent
func NewValidatePushEvent(repositoryBranchToConsider string, statusMessageOnInvalidBranch string) *ValidatePushEvent {
	handler := new(ValidatePushEvent)
	handler.repositoryBranchToConsider = repositoryBranchToConsider
	handler.statusMessageOnInvalidBranch = statusMessageOnInvalidBranch
	return handler
}
