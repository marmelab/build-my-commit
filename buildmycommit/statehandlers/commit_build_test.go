package statehandlers

import (
	"errors"
	"github.com/marmelab/buildmycommit/states"
	"testing"
)

func TestCommitOutputShouldReturnStateInternalServerErrorWhenGitThrowsError(t *testing.T) {
	// Mock the executeGitWithContext func
	executeGitWithContext := func(command string, repositoryPath string, arguments ...string) (string, error) {
		return "", errors.New("42")
	}

	state := states.State{RepositoryPath: "repository-path"}

	handler := NewCommitBuild(executeGitWithContext, "commit-message")

	newState, _ := handler.Handle(state)

	if newState != states.InternalServerError {
		t.Errorf("HandleStateCommitOutput returned %v instead of %v", newState, states.InternalServerError)
	}
}

func TestCommitOutputShouldReturnStateVerifyRemoteHash(t *testing.T) {
	// Mock the executeGitWithContext func
	executeGitWithContext := func(command string, repositoryPath string, arguments ...string) (string, error) {
		return "", nil
	}

	state := states.State{RepositoryPath: "repository-path"}

	handler := NewCommitBuild(executeGitWithContext, "commit-message")

	newState, _ := handler.Handle(state)

	if newState != states.CheckRemoteHash {
		t.Errorf("HandleStateCommitOutput returned %v instead of %v", newState, states.CheckRemoteHash)
	}
}
