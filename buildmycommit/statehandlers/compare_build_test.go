package statehandlers

import (
	"errors"
	"github.com/marmelab/buildmycommit/states"
	"net/http"
	"testing"
)

func TestCompareOutputShouldReturnStateInternalServerErrorWhenGitThrowsError(t *testing.T) {
	// Mock the executeGitWithContext func
	executeGitWithContext := func(command string, repositoryPath string, arguments ...string) (string, error) {
		return "", errors.New("42")
	}

	state := states.State{RepositoryPath: "repository-path"}

	handler := NewCompareBuild(executeGitWithContext, "")

	newState, _ := handler.Handle(state)

	if newState != states.InternalServerError {
		t.Errorf("HandleStateCompareOutput returned %v instead of %v", newState, states.InternalServerError)
	}
}

func TestCompareOutputShouldReturnStateEndRequestIfRemoteOutputIsUpToDate(t *testing.T) {
	// Mock the executeGitWithContext func
	executeGitWithContext := func(command string, repositoryPath string, arguments ...string) (string, error) {
		return "", nil
	}

	state := states.State{RepositoryPath: "repository-path"}

	handler := NewCompareBuild(executeGitWithContext, "")

	newState, _ := handler.Handle(state)

	if newState != states.EndRequest {
		t.Errorf("HandleStateCompareOutput returned %v instead of %v", newState, states.EndRequest)
	}
}

func TestCompareOutputShouldSetStateStatusToOKIfRemoteOutputIsUpToDate(t *testing.T) {
	// Mock the executeGitWithContext func
	executeGitWithContext := func(command string, repositoryPath string, arguments ...string) (string, error) {
		return "", nil
	}

	state := states.State{RepositoryPath: "repository-path"}

	handler := NewCompareBuild(executeGitWithContext, "")

	_, newStatePayload := handler.Handle(state)

	if newStatePayload.Status != http.StatusOK {
		t.Errorf("HandleStateCompareOutput returned %v instead of %v", newStatePayload.Status, http.StatusOK)
	}
}

func TestCompareOutputShouldSetStateStatusMessageIfRemoteOutputIsUpToDate(t *testing.T) {
	// Mock the executeGitWithContext func
	executeGitWithContext := func(command string, repositoryPath string, arguments ...string) (string, error) {
		return "", nil
	}

	state := states.State{RepositoryPath: "repository-path"}

	handler := NewCompareBuild(executeGitWithContext, "message")

	_, newStatePayload := handler.Handle(state)

	if newStatePayload.StatusMessage != "message" {
		t.Errorf("HandleStateCompareOutput returned %v instead of %v", newStatePayload.StatusMessage, "message")
	}
}

func TestCompareOutputShouldReturnStateEndRequestIfRemoteOutputIsNotUpToDate(t *testing.T) {
	// Mock the executeGitWithContext func
	executeGitWithContext := func(command string, repositoryPath string, arguments ...string) (string, error) {
		return "status-indicating-diff-with-remote", nil
	}

	state := states.State{RepositoryPath: "repository-path"}

	handler := NewCompareBuild(executeGitWithContext, "")

	newState, _ := handler.Handle(state)

	if newState != states.CommitBuild {
		t.Errorf("HandleStateCompareOutput returned %v instead of %v", newState, states.CommitBuild)
	}
}
