package statehandlers

import (
	"github.com/marmelab/buildmycommit/states"
	"github.com/marmelab/buildmycommit/git"
	"net/http"
	"strings"
)

// ValidateCommit is the StateHandler in charge of the ValidateCommit state
type ValidateCommit struct {
	executeGitInContext                     git.ExecuteGitWithContext
	commitMessage                           string
	statusMessageWhenAutomaticBuildDetected string
}

// Handle the ValidateCommit state
func (stateHandler ValidateCommit) Handle(state states.State) (int, states.State) {
	// Get the last commit message without pretty formatting
	output, err := stateHandler.executeGitInContext(
		"log",
		state.RepositoryPath,
		"-1",
		"--pretty=%B")

	if err != nil {
		return states.InternalServerError, state
	}

	// If the last commit message is our standard commit message, we end here
	if strings.Contains(output, stateHandler.commitMessage) {
		state.Status = http.StatusOK
		state.StatusMessage = stateHandler.statusMessageWhenAutomaticBuildDetected
		return states.EndRequest, state
	}

	return states.BuildDocker, state
}

// NewValidateCommit is a constructor like factory which returns a ValidateCommit
func NewValidateCommit(executeGitInContext git.ExecuteGitWithContext, commitMessage string, statusMessageWhenAutomaticBuildDetected string) *ValidateCommit {
	handler := new(ValidateCommit)
	handler.executeGitInContext = executeGitInContext
	handler.commitMessage = commitMessage
	handler.statusMessageWhenAutomaticBuildDetected = statusMessageWhenAutomaticBuildDetected

	return handler
}
