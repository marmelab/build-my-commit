package statehandlers

import (
	"errors"
	"github.com/marmelab/buildmycommit/states"
	"net/http"
	"testing"
)

func TestVerifyRemoteHashShouldReturnStateInternalServerErrorWhenGitThrowsError(t *testing.T) {
	// Mock the gitWithContext func
	gitGitWithContext := func(command string, contextPath string, arguments ...string) (string, error) {
		return "", errors.New("42")
	}

	state := states.State{}

	handler := NewCheckRemoteHash(gitGitWithContext, "")

	newState, _ := handler.Handle(state)

	if newState != states.InternalServerError {
		t.Errorf("HandleStateVerifyRemoteHash returned %v instead of %v", newState, states.InternalServerError)
	}
}

func TestVerifyRemoteHashShouldReturnStateEndRequestIfRemoteOutputHasChanged(t *testing.T) {
	// Mock the gitWithContext func
	gitGitWithContext := func(command string, contextPath string, arguments ...string) (string, error) {
		return "has-changed", nil
	}

	pushEvent := states.PushEvent{}
	pushEvent.HeadCommit.ID = "test"
	state := states.State{RepositoryPath: "repository-path", PushEvent: pushEvent}

	handler := NewCheckRemoteHash(gitGitWithContext, "")

	newState, _ := handler.Handle(state)

	if newState != states.EndRequest {
		t.Errorf("HandleStateVerifyRemoteHash returned %v instead of %v", newState, states.EndRequest)
	}
}

func TestVerifyRemoteHashShouldSetStateStatusToOKRemoteOutputHasChanged(t *testing.T) {
	// Mock the gitWithContext func
	gitGitWithContext := func(command string, contextPath string, arguments ...string) (string, error) {
		return "has-changed", nil
	}

	pushEvent := states.PushEvent{}
	pushEvent.HeadCommit.ID = "test"
	state := states.State{RepositoryPath: "repository-path", PushEvent: pushEvent}

	handler := NewCheckRemoteHash(gitGitWithContext, "")

	_, newStatePayload := handler.Handle(state)

	if newStatePayload.Status != http.StatusOK {
		t.Errorf("HandleStateVerifyRemoteHash returned %v instead of %v", newStatePayload.Status, http.StatusOK)
	}
}

func TestVerifyRemoteHashShouldSetStateStatusMessageRemoteOutputHasChanged(t *testing.T) {
	// Mock the gitWithContext func
	gitGitWithContext := func(command string, contextPath string, arguments ...string) (string, error) {
		return "has-changed", nil
	}

	pushEvent := states.PushEvent{}
	pushEvent.HeadCommit.ID = "test"
	state := states.State{RepositoryPath: "repository-path", PushEvent: pushEvent}

	handler := NewCheckRemoteHash(gitGitWithContext, "message")

	_, newStatePayload := handler.Handle(state)

	if newStatePayload.StatusMessage != "message" {
		t.Errorf("HandleStateVerifyRemoteHash returned %v instead of %v", newStatePayload.StatusMessage, "message")
	}
}

func TestVerifyRemoteHashShouldReturnStateEndRequestIfRemoteOutputHasNotChanged(t *testing.T) {
	// Mock the gitWithContext func
	gitGitWithContext := func(command string, contextPath string, arguments ...string) (string, error) {
		return "test", nil
	}

	pushEvent := states.PushEvent{}
	pushEvent.HeadCommit.ID = "test"
	state := states.State{RepositoryPath: "repository-path", PushEvent: pushEvent}

	handler := NewCheckRemoteHash(gitGitWithContext, "")

	newState, _ := handler.Handle(state)

	if newState != states.PushBuild {
		t.Errorf("HandleStateVerifyRemoteHash returned %v instead of %v", newState, states.PushBuild)
	}
}
