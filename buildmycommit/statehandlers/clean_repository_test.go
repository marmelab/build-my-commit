package statehandlers

import (
	"errors"
	"github.com/marmelab/buildmycommit/states"
	"testing"
)

func TestCleanGitShouldReturnStateInternalServerErrorWhenOsRemoveAllThrowsError(t *testing.T) {
	// Mock the removeDirectory func
	removeDirectory := func(path string) error {
		return errors.New("42")
	}

	state := states.State{}

	handler := NewCleanRepository(removeDirectory)

	newState, _ := handler.Handle(state)

	if newState != states.InternalServerError {
		t.Errorf("HandleStateCleanGit returned %v instead of %v", newState, states.InternalServerError)
	}
}

func TestCleanGitShouldReturnStateEndRequest(t *testing.T) {
	// Mock the removeDirectory func
	removeDirectory := func(path string) error {
		return nil
	}

	state := states.State{}

	handler := NewCleanRepository(removeDirectory)

	newState, _ := handler.Handle(state)

	if newState != states.EndRequest {
		t.Errorf("HandleStateCleanGit returned %v instead of %v", newState, states.RunDocker)
	}
}

func TestCleanGitShouldSetShouldCleanGitToFalse(t *testing.T) {
	// Mock the removeDirectory func
	removeDirectory := func(path string) error {
		return nil
	}

	state := states.State{}

	handler := NewCleanRepository(removeDirectory)

	_, newStatePayload := handler.Handle(state)

	if newStatePayload.ShouldCleanGit {
		t.Errorf("HandleStateCleanGit should have set ShouldCleanGit to false")
	}
}
