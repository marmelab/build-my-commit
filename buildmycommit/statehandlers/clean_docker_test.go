package statehandlers

import (
	"errors"
	"github.com/marmelab/buildmycommit/states"
	"testing"
)

func TestCleanDockerShouldReturnStateInternalServerErrorWhenDockerThrowsError(t *testing.T) {
	// Mock the docker func
	executeDocker := func(command string, arguments ...string) error {
		return errors.New("42")
	}

	state := states.State{}

	handler := NewCleanDocker(executeDocker)

	newState, _ := handler.Handle(state)

	if newState != states.InternalServerError {
		t.Errorf("HandleStateCleanDocker returned %v instead of %v", newState, states.InternalServerError)
	}
}

func TestCleanDockerShouldReturnStateEndRequest(t *testing.T) {
	// Mock the docker func
	executeDocker := func(command string, arguments ...string) error {
		return nil
	}

	state := states.State{}

	handler := NewCleanDocker(executeDocker)

	newState, _ := handler.Handle(state)

	if newState != states.EndRequest {
		t.Errorf("HandleStateCleanDocker returned %v instead of %v", newState, states.RunDocker)
	}
}

func TestCleanDockerShouldSetShouldCleanDockerToFalse(t *testing.T) {
	// Mock the docker func
	executeDocker := func(command string, arguments ...string) error {
		return nil
	}

	state := states.State{}

	handler := NewCleanDocker(executeDocker)

	_, newStatePayload := handler.Handle(state)

	if newStatePayload.ShouldCleanDocker {
		t.Errorf("HandleStateCleanDocker should have set ShouldCleanDocker to false")
	}
}
