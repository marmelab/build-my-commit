package statehandlers

import (
	"errors"
	"github.com/marmelab/buildmycommit/states"
	"net/http"
	"testing"
)

func TestValidateDockerFileShouldReturnStateInternalServerErrorWhenGitThrowsError(t *testing.T) {
	// Mock the gitGitWithContext func
	executeGitInContext := func(command string, contextPath string, arguments ...string) (string, error) {
		return "", errors.New("42")
	}

	state := states.State{RepositoryPath: "repository-path"}

	handler := NewValidateCommit(executeGitInContext, "", "")

	newState, _ := handler.Handle(state)

	if newState != states.InternalServerError {
		t.Errorf("HandleStateValidateDockerFile returned %v instead of %v", newState, states.InternalServerError)
	}
}

func TestValidateDockerFileShouldReturnStateEndRequestIfRemoteOutputIsUpToDate(t *testing.T) {
	// Mock the gitGitWithContext func
	executeGitInContext := func(command string, contextPath string, arguments ...string) (string, error) {
		return "commitMessage", nil
	}

	state := states.State{RepositoryPath: "repository-path"}

	handler := NewValidateCommit(executeGitInContext, "commitMessage", "")

	newState, _ := handler.Handle(state)

	if newState != states.EndRequest {
		t.Errorf("HandleStateValidateDockerFile returned %v instead of %v", newState, states.EndRequest)
	}
}

func TestValidateDockerFileShouldSetStateStatusToOKIfRemoteOutputIsUpToDate(t *testing.T) {
	// Mock the gitGitWithContext func
	executeGitInContext := func(command string, contextPath string, arguments ...string) (string, error) {
		return "commitMessage", nil
	}

	state := states.State{RepositoryPath: "repository-path"}

	handler := NewValidateCommit(executeGitInContext, "commitMessage", "")

	_, newStatePayload := handler.Handle(state)

	if newStatePayload.Status != http.StatusOK {
		t.Errorf("HandleStateValidateDockerFile returned %v instead of %v", newStatePayload.Status, http.StatusOK)
	}
}

func TestValidateDockerFileShouldSetStateStatusMessageIfRemoteOutputIsUpToDate(t *testing.T) {
	// Mock the gitGitWithContext func
	executeGitInContext := func(command string, contextPath string, arguments ...string) (string, error) {
		// Return an empty string because we want no diff between local and remote
		return "commitMessage", nil
	}

	state := states.State{RepositoryPath: "repository-path"}

	handler := NewValidateCommit(executeGitInContext, "commitMessage", "message")

	_, newStatePayload := handler.Handle(state)

	if newStatePayload.StatusMessage != "message" {
		t.Errorf("HandleStateValidateDockerFile returned %v instead of %v", newStatePayload.StatusMessage, "message")
	}
}

func TestValidateDockerFileShouldReturnStateEndRequestIfRemoteOutputIsNotUpToDate(t *testing.T) {
	// Mock the gitGitWithContext func
	executeGitInContext := func(command string, contextPath string, arguments ...string) (string, error) {
		return "anotherCommitMessage", nil
	}

	state := states.State{RepositoryPath: "repository-path"}

	handler := NewValidateCommit(executeGitInContext, "commitMessage", "")

	newState, _ := handler.Handle(state)

	if newState != states.BuildDocker {
		t.Errorf("HandleStateValidateDockerFile returned %v instead of %v", newState, states.BuildDocker)
	}
}
