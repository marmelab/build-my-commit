package main

import (
	"errors"
	"testing"
)

func TestHandleStateCloneRepositoryShouldReturnStateEndRequestWithErrorWhenGitThrowsError(t *testing.T) {
	// Mock the os.RemoveAll func
	defer Patch(&git, func(command string, arguments ...string) (string, error) {
		return "", errors.New("42")
	}).Restore()

	state := State{}

	newState, _, err := HandleStateCloneRepository(state)

	if newState != StateEndRequest {
		t.Errorf("HandleStateCloneRepository returned %v instead of %v", newState, StateEndRequest)
	}

	if err == nil {
		t.Error("HandleStateCloneRepository should have returned the error")
	}
}

func TestHandleStateCloneRepositoryShouldStateValidateRepository(t *testing.T) {
	// Mock the os.RemoveAll func
	defer Patch(&git, func(command string, arguments ...string) (string, error) {
		return "", nil
	}).Restore()

	state := State{}

	newState, _, err := HandleStateCloneRepository(state)

	if err != nil {
		t.Errorf("HandleStateCloneRepository returned the error %q", err)
	}

	if newState != StateValidateRepository {
		t.Errorf("HandleStateCloneRepository returned %v instead of %v", newState, StateValidateRepository)
	}
}

func TestHandleStateCloneRepositoryShouldSetShouldCleanGitToTrue(t *testing.T) {
	// Mock the os.RemoveAll func
	defer Patch(&git, func(command string, arguments ...string) (string, error) {
		return "", nil
	}).Restore()

	state := State{}

	_, newStatePayload, err := HandleStateCloneRepository(state)

	if err != nil {
		t.Errorf("HandleStateCleanDocker returned the error %q", err)
	}

	if !newStatePayload.ShouldCleanGit {
		t.Errorf("HandleStateCleanDocker should have set ShouldCleanGit to true")
	}
}
