package statehandlers

import (
	"errors"
	"github.com/marmelab/buildmycommit/states"
	"testing"
)

func TestPushBuildShouldReturnStateInternalServerErrorWhenGitThrowsError(t *testing.T) {
	// Mock the gitWithContext func
	executeGitInContext := func(command string, contextPath string, arguments ...string) (string, error) {
		return "", errors.New("42")
	}

	state := states.State{}

	handler := NewPushBuild(executeGitInContext)

	newState, _ := handler.Handle(state)

	if newState != states.InternalServerError {
		t.Errorf("HandleStatePushBuild returned %v instead of %v", newState, states.InternalServerError)
	}
}

func TestPushBuildShouldReturnStateVerifyRemoteHash(t *testing.T) {
	// Mock the gitWithContext func
	executeGitInContext := func(command string, contextPath string, arguments ...string) (string, error) {
		return "", nil
	}

	state := states.State{RepositoryPath: "repository-path"}

	handler := NewPushBuild(executeGitInContext)

	newState, _ := handler.Handle(state)

	if newState != states.EndRequest {
		t.Errorf("HandleStatePushBuild returned %v instead of %v", newState, states.EndRequest)
	}
}
