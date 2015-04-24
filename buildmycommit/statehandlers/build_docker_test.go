package statehandlers

import (
	"errors"
	"github.com/marmelab/buildmycommit/states"
	"testing"
)

func TestBuildDockerShouldReturnStateInternalServerErrorWhenDockerThrowsError(t *testing.T) {
	// Mock the docker func
	executeDocker := func(command string, arguments ...string) error {
		return errors.New("42")
	}

	state := states.State{}

	handler := NewBuildDocker(executeDocker, "dockerFilePath")

	newState, _ := handler.Handle(state)

	if newState != states.InternalServerError {
		t.Errorf("HandleStateBuildDocker returned %v instead of %v", newState, states.InternalServerError)
	}
}

func TestBuildDockerShouldReturnStateRunDocker(t *testing.T) {
	// Mock the exists func
	executeDocker := func(command string, arguments ...string) error {
		return nil
	}

	state := states.State{}

	handler := NewBuildDocker(executeDocker, "dockerFilePath")

	newState, _ := handler.Handle(state)

	if newState != states.RunDocker {
		t.Errorf("HandleStateBuildDocker returned %v instead of %v", newState, states.RunDocker)
	}
}

func TestBuildDockerShouldSetShouldCleanDockerToTrue(t *testing.T) {
	// Mock the exists func
	executeDocker := func(command string, arguments ...string) error {
		return nil
	}

	state := states.State{}

	handler := NewBuildDocker(executeDocker, "dockerFilePath")

	_, newStatePayload := handler.Handle(state)

	if !newStatePayload.ShouldCleanDocker {
		t.Errorf("HandleStateBuildDocker should have set ShouldCleanDocker to true")
	}
}
