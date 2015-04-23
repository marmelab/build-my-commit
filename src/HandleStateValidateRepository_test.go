package main

import (
	"errors"
	"net/http"
	"testing"
)

func TestHandleStateValidateRepositoryShouldReturnStateEndRequestWithErrorWhenExistsThrowsError(t *testing.T) {
	// Mock the gitWithContext func
	defer Patch(&exists, func(path string) (bool, error) {
		return false, errors.New("42")
	}).Restore()

	state := State{RepositoryPath: "repository-path"}

	newState, _, err := HandleStateValidateRepository(state)

	if newState != StateEndRequest {
		t.Errorf("HandleStateValidateRepository returned %v instead of %v", newState, StateEndRequest)
	}

	if err == nil {
		t.Error("HandleStateValidateRepository should have returned the error")
	}
}

func TestHandleStateValidateRepositoryShouldReturnStateEndRequestWithErrorWhenGitThrowsError(t *testing.T) {
	// Mock the gitWithContext func
	defer Patch(&exists, func(path string) (bool, error) {
		return true, nil
	}).Restore()

	// Mock the gitWithContext func
	defer Patch(&gitWithContext, func(command string, contextPath string, arguments ...string) (string, error) {
		return "", errors.New("42")
	}).Restore()

	state := State{RepositoryPath: "repository-path"}

	newState, _, err := HandleStateValidateRepository(state)

	if newState != StateEndRequest {
		t.Errorf("HandleStateValidateRepository returned %v instead of %v", newState, StateEndRequest)
	}

	if err == nil {
		t.Error("HandleStateValidateRepository should have returned the error")
	}
}

func TestHandleStateValidateRepositoryShouldReturnStateEndRequestIfExistsReturnFalse(t *testing.T) {
	// Mock the exists func
	defer Patch(&exists, func(path string) (bool, error) {
		return false, nil
	}).Restore()

	state := State{RepositoryPath: "repository-path"}

	newState, _, err := HandleStateValidateRepository(state)

	if err != nil {
		t.Errorf("HandleStateValidateRepository returned the error %q", err)
	}

	if newState != StateEndRequest {
		t.Errorf("HandleStateValidateRepository returned %v instead of %v", newState, StateEndRequest)
	}
}

func TestHandleStateValidateRepositoryShouldReturnStateEndRequestIfRemoteOutputIsUpToDate(t *testing.T) {
	// Mock the exists func
	defer Patch(&exists, func(path string) (bool, error) {
		return true, nil
	}).Restore()

	// Mock the gitWithContext func
	defer Patch(&gitWithContext, func(command string, contextPath string, arguments ...string) (string, error) {
		return CommitMessage, nil
	}).Restore()

	state := State{RepositoryPath: "repository-path"}

	newState, _, err := HandleStateValidateRepository(state)

	if err != nil {
		t.Errorf("HandleStateValidateRepository returned the error %q", err)
	}

	if newState != StateEndRequest {
		t.Errorf("HandleStateValidateRepository returned %v instead of %v", newState, StateEndRequest)
	}
}

func TestHandleStateValidateRepositoryShouldSetStateStatusToOKIfRemoteOutputIsUpToDate(t *testing.T) {
	// Mock the exists func
	defer Patch(&exists, func(path string) (bool, error) {
		return true, nil
	}).Restore()

	// Mock the gitWithContext func
	defer Patch(&gitWithContext, func(command string, contextPath string, arguments ...string) (string, error) {
		return CommitMessage, nil
	}).Restore()

	state := State{RepositoryPath: "repository-path"}

	_, newStatePayload, err := HandleStateValidateRepository(state)

	if err != nil {
		t.Errorf("HandleStateValidateRepository returned the error %q", err)
	}

	if newStatePayload.Status != http.StatusOK {
		t.Errorf("HandleStateValidateRepository returned %v instead of %v", newStatePayload.Status, http.StatusOK)
	}
}

func TestHandleStateValidateRepositoryShouldSetStateStatusMessageIfRemoteOutputIsUpToDate(t *testing.T) {
	// Mock the exists func
	defer Patch(&exists, func(path string) (bool, error) {
		return true, nil
	}).Restore()

	// Mock the gitWithContext func
	defer Patch(&gitWithContext, func(command string, contextPath string, arguments ...string) (string, error) {
		// Return an empty string because we want no diff between local and remote
		return CommitMessage, nil
	}).Restore()

	state := State{RepositoryPath: "repository-path"}

	_, newStatePayload, err := HandleStateValidateRepository(state)

	if err != nil {
		t.Errorf("HandleStateValidateRepository returned the error %q", err)
	}

	if newStatePayload.StatusMessage != StatusMessageForStateValidateRepositoryCommitDetected {
		t.Errorf("HandleStateValidateRepository returned %v instead of %v", newStatePayload.StatusMessage, StatusMessageForStateValidateRepositoryCommitDetected)
	}
}

func TestHandleStateValidateRepositoryShouldReturnStateEndRequestIfRemoteOutputIsNotUpToDate(t *testing.T) {
	// Mock the exists func
	defer Patch(&exists, func(path string) (bool, error) {
		return true, nil
	}).Restore()

	// Mock the gitWithContext func
	defer Patch(&gitWithContext, func(command string, contextPath string, arguments ...string) (string, error) {
		return "", nil
	}).Restore()

	state := State{RepositoryPath: "repository-path"}

	newState, _, err := HandleStateValidateRepository(state)

	if err != nil {
		t.Errorf("HandleStateValidateRepository returned the error %q", err)
	}

	if newState != StateBuildDocker {
		t.Errorf("HandleStateValidateRepository returned %v instead of %v", newState, StateBuildDocker)
	}
}
