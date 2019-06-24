package statehandlers

import (
	"errors"
	"testing"

	"github.com/marmelab/buildmycommit/states"
)

func TestCommitOutputShouldReturnStateInternalServerErrorWhenGitThrowsErrorOnAdd(t *testing.T) {
	// Mock the executeGitWithContext func
	executeGitWithContext := func(command string, repositoryPath string, arguments ...string) (string, error) {
		if command == "add" {
			return "", errors.New("42")
		}
		return "", nil
	}

	state := states.State{RepositoryPath: "repository-path"}

	handler := NewCommitBuild(executeGitWithContext, "commit-message")

	newState, _ := handler.Handle(state)

	if newState != states.InternalServerError {
		t.Errorf("HandleStateCommitOutput returned %v instead of %v", newState, states.InternalServerError)
	}
}

func TestCommitOutputShouldReturnStateInternalServerErrorWhenGitThrowsErrorOnCommit(t *testing.T) {
	// Mock the executeGitWithContext func
	executeGitWithContext := func(command string, repositoryPath string, arguments ...string) (string, error) {
		if command == "commit" {
			return "", errors.New("42")
		}
		return "", nil
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
