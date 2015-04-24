package statehandlers

import (
	"fmt"
	"github.com/marmelab/buildmycommit/states"
	"github.com/marmelab/buildmycommit/git"
)

// CommitBuild is the StateHandler in charge of the CommitBuild state
type CommitBuild struct {
	executeGitInContext git.ExecuteGitWithContext
	commitMessage       string
}

// Handle the CommitBuild state
func (stateHandler CommitBuild) Handle(state states.State) (int, states.State) {
	// Add files as it may be the first build
	_, err := stateHandler.executeGitInContext(
		"add",
		state.RepositoryPath,
		"build")

	if err != nil {
		return states.InternalServerError, state
	}

	// Commit files
	message := stateHandler.commitMessage + " " + state.PushEvent.HeadCommit.ID

	if err != nil {
		return states.InternalServerError, state
	}

	_, err = stateHandler.executeGitInContext(
		"commit",
		state.RepositoryPath,
		fmt.Sprintf("--message=%v", message),
		".")

	if err != nil {
		return states.InternalServerError, state
	}

	return states.CheckRemoteHash, state
}

// NewCommitBuild is a constructor like factory which returns a CommitBuild
func NewCommitBuild(executeGitInContext git.ExecuteGitWithContext, commitMessage string) *CommitBuild {
	handler := new(CommitBuild)
	handler.executeGitInContext = executeGitInContext
	handler.commitMessage = commitMessage

	return handler
}
