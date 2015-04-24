package statehandlers

import (
	"errors"
	"github.com/marmelab/buildmycommit/states"
	"net/http"
	"testing"
)

func TestValidateDockerFileShouldReturnStateInternalServerErrorWhenExistsThrowsError(t *testing.T) {
	// Mock the exists func
	exists := func(path string) (bool, error) {
		return false, errors.New("42")
	}

	state := states.State{RepositoryPath: "repository-path"}

	handler := NewValidateDockerFile(exists, "", "")

	newState, _ := handler.Handle(state)

	if newState != states.InternalServerError {
		t.Errorf("HandleStateValidateDockerFile returned %v instead of %v", newState, states.InternalServerError)
	}
}

func TestValidateDockerFileShouldReturnStateEndRequestIfExistsReturnFalse(t *testing.T) {
	// Mock the exists func
	exists := func(path string) (bool, error) {
		return false, nil
	}

	state := states.State{RepositoryPath: "repository-path"}

	handler := NewValidateDockerFile(exists, "dockerFilePath", "")

	newState, _ := handler.Handle(state)

	if newState != states.EndRequest {
		t.Errorf("HandleStateValidateDockerFile returned %v instead of %v", newState, states.EndRequest)
	}
}

func TestValidateDockerFileShouldSetStateStatusToOKIfExistsReturnFalse(t *testing.T) {
	// Mock the exists func
	exists := func(path string) (bool, error) {
		return false, nil
	}

	state := states.State{RepositoryPath: "repository-path"}

	handler := NewValidateDockerFile(exists, "dockerFilePath", "")

	_, newStatePayload := handler.Handle(state)

	if newStatePayload.Status != http.StatusOK {
		t.Errorf("HandleStateValidateDockerFile returned %v instead of %v", newStatePayload.Status, http.StatusOK)
	}
}

func TestValidateDockerFileShouldSetStateStatusMessageIfExistsReturnFalse(t *testing.T) {
	// Mock the exists func
	exists := func(path string) (bool, error) {
		return false, nil
	}

	state := states.State{RepositoryPath: "repository-path"}

	handler := NewValidateDockerFile(exists, "dockerFilePath", "message")

	_, newStatePayload := handler.Handle(state)

	if newStatePayload.StatusMessage != "message" {
		t.Errorf("HandleStateValidateDockerFile returned %v instead of %v", newStatePayload.StatusMessage, "message")
	}
}

func TestValidateDockerFileShouldReturnStateValidateCommitIfExistsReturnTrue(t *testing.T) {
	// Mock the exists func
	exists := func(path string) (bool, error) {
		return true, nil
	}

	state := states.State{RepositoryPath: "repository-path"}

	handler := NewValidateDockerFile(exists, "dockerFilePath", "")

	newState, _ := handler.Handle(state)

	if newState != states.ValidateCommit {
		t.Errorf("HandleStateValidateDockerFile returned %v instead of %v", newState, states.BuildDocker)
	}
}
