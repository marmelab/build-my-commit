package statehandlers

import (
	"github.com/marmelab/buildmycommit/states"
	"github.com/marmelab/buildmycommit/git"
)

// PushBuild is the StateHandler in charge of the PushBuild state
type PushBuild struct {
	executeGitInContext git.ExecuteGitWithContext
}

// Handle the PushBuild state
func (stateHandler PushBuild) Handle(state states.State) (int, states.State) {
	_, err := stateHandler.executeGitInContext(
		"push",
		state.RepositoryPath)

	if err != nil {
		return states.InternalServerError, state
	}

	return states.EndRequest, state
}

// NewPushBuild is a constructor like factory which returns a PushBuild
func NewPushBuild(executeGitInContext git.ExecuteGitWithContext) *PushBuild {
	handler := new(PushBuild)
	handler.executeGitInContext = executeGitInContext

	return handler
}
