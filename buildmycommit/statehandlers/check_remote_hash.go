package statehandlers

import (
	"github.com/marmelab/buildmycommit/states"
	"github.com/marmelab/buildmycommit/git"
	"net/http"
)

// CheckRemoteHash is the StateHandler in charge of the CheckRemoteHash state
type CheckRemoteHash struct {
	executeGitWithContext                 git.ExecuteGitWithContext
	statusMessageWhenRemoteHasBeenUpdated string
}

// Handle the CheckRemoteHash state
func (stateHandler CheckRemoteHash) Handle(state states.State) (int, states.State) {
	lastHash, err := stateHandler.executeGitWithContext(
		"rev-parse",
		state.RepositoryPath,
		"origin/master")

	if err != nil {
		return states.InternalServerError, state
	}

	// Repository hasn't been updated if we get the same hash as pushEvent
	// If not, we just discard this build as another one is probably running in parallel
	if lastHash == state.PushEvent.HeadCommit.ID {
		return states.PushBuild, state
	}

	state.Status = http.StatusOK
	state.StatusMessage = stateHandler.statusMessageWhenRemoteHasBeenUpdated
	return states.EndRequest, state
}

// NewCheckRemoteHash is a constructor like factory which returns a CheckRemoteHash
func NewCheckRemoteHash(executeGitWithContext git.ExecuteGitWithContext, statusMessageWhenRemoteHasBeenUpdated string) *CheckRemoteHash {
	handler := new(CheckRemoteHash)
	handler.executeGitWithContext = executeGitWithContext
	handler.statusMessageWhenRemoteHasBeenUpdated = statusMessageWhenRemoteHasBeenUpdated

	return handler
}
