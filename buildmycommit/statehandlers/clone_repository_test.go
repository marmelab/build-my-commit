package statehandlers

import (
	"errors"
	"github.com/marmelab/buildmycommit/states"
	"testing"
)

func TestCloneRepositoryShouldReturnStateInternalServerErrorWhenGitThrowsError(t *testing.T) {
	// Mock the executeGit func
	executeGit := func(command string, arguments ...string) (string, error) {
		return "", errors.New("42")
	}

	state := states.State{}

	handler := NewCloneRepository(executeGit)

	newState, _ := handler.Handle(state)

	if newState != states.InternalServerError {
		t.Errorf("HandleStateCloneRepository returned %v instead of %v", newState, states.InternalServerError)
	}
}

func TestCloneRepositoryShouldStateValidateDockerFile(t *testing.T) {
	// Mock the executeGit func
	executeGit := func(command string, arguments ...string) (string, error) {
		return "", nil
	}

	state := states.State{}

	handler := NewCloneRepository(executeGit)

	newState, _ := handler.Handle(state)

	if newState != states.ValidateDockerFile {
		t.Errorf("HandleStateCloneRepository returned %v instead of %v", newState, states.ValidateDockerFile)
	}
}

func TestCloneRepositoryShouldSetShouldCleanGitToTrue(t *testing.T) {
	// Mock the executeGit func
	executeGit := func(command string, arguments ...string) (string, error) {
		return "", nil
	}

	state := states.State{}

	handler := NewCloneRepository(executeGit)

	_, newStatePayload := handler.Handle(state)

	if !newStatePayload.ShouldCleanGit {
		t.Errorf("HandleStateCleanDocker should have set ShouldCleanGit to true")
	}
}
