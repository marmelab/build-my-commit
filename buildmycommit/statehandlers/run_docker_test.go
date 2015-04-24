package statehandlers

import (
	"errors"
	"github.com/marmelab/buildmycommit/states"
	"testing"
)

func TestRunDockerShouldReturnStateInternalServerErrorWhenMakeAbsolutePathThrowsError(t *testing.T) {
	// Mock the filepath.Abs func
	makeAbsolutePath := func(path string) (string, error) {
		return "", errors.New("42")
	}

	// Mock the docker func
	executeDocker := func(command string, arguments ...string) error {
		return nil
	}

	state := states.State{}

	handler := NewRunDocker(executeDocker, makeAbsolutePath)

	newState, _ := handler.Handle(state)

	if newState != states.InternalServerError {
		t.Errorf("HandleStateRunDocker returned %v instead of %v", newState, states.InternalServerError)
	}
}

func TestRunDockerShouldReturnStateInternalServerErrorWhenDockerThrowsError(t *testing.T) {
	// Mock the filepath.Abs func
	makeAbsolutePath := func(path string) (string, error) {
		return "", nil
	}

	// Mock the docker func
	executeDocker := func(command string, arguments ...string) error {
		return errors.New("42")
	}

	state := states.State{}

	handler := NewRunDocker(executeDocker, makeAbsolutePath)

	newState, _ := handler.Handle(state)

	if newState != states.InternalServerError {
		t.Errorf("HandleStateRunDocker returned %v instead of %v", newState, states.InternalServerError)
	}
}

func TestRunDockerShouldReturnStateRunDocker(t *testing.T) {
	// Mock the filepath.Abs func
	makeAbsolutePath := func(path string) (string, error) {
		return "", nil
	}

	// Mock the docker func
	executeDocker := func(command string, arguments ...string) error {
		return nil
	}

	state := states.State{}

	handler := NewRunDocker(executeDocker, makeAbsolutePath)

	newState, _ := handler.Handle(state)

	if newState != states.CompareBuild {
		t.Errorf("HandleStateRunDocker returned %v instead of %v", newState, states.CompareBuild)
	}
}
