package statehandlers

import (
	"github.com/marmelab/buildmycommit/states"
	"github.com/marmelab/buildmycommit/git"
	"net/http"
)

// CompareBuild is the StateHandler in charge of the CompareBuild state
type CompareBuild struct {
	executeGitInContext               git.ExecuteGitWithContext
	statusMessageWhenRemoteIsUpToDate string
}

// Handle the CompareBuild state
func (stateHandler CompareBuild) Handle(state states.State) (int, states.State) {
	output, err := stateHandler.executeGitInContext(
		"status",
		state.RepositoryPath,
		"--porcelain") // This will make git status return a machine readable output without pretty formatting

	if err != nil {
		return states.InternalServerError, state
	}

	// 7. Commit & push the output if necessary
	if output != "" && len(output) > 0 {
		return states.CommitBuild, state
	}

	state.Status = http.StatusOK
	state.StatusMessage = stateHandler.statusMessageWhenRemoteIsUpToDate
	return states.EndRequest, state
}

// NewCompareBuild is a constructor like factory which returns a CompareBuild
func NewCompareBuild(executeGitInContext git.ExecuteGitWithContext, statusMessageWhenRemoteIsUpToDate string) *CompareBuild {
	handler := new(CompareBuild)
	handler.executeGitInContext = executeGitInContext
	handler.statusMessageWhenRemoteIsUpToDate = statusMessageWhenRemoteIsUpToDate

	return handler
}
