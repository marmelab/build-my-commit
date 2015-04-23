package main

import (
	"errors"
	"testing"
)

func TestHandleStateCleanGitShouldReturnStateEndRequestWithErrorWhenOsRemoveAllThrowsError(t *testing.T) {
	// Mock the os.RemoveAll func
	defer Patch(&osRemoveAll, func(path string) error {
		return errors.New("42")
	}).Restore()

	state := State{}

	newState, _, err := HandleStateCleanGit(state)

	if newState != StateEndRequest {
		t.Errorf("HandleStateCleanGit returned %v instead of %v", newState, StateEndRequest)
	}

	if err == nil {
		t.Error("HandleStateCleanGit should have returned the error")
	}
}

func TestHandleStateCleanGitShouldReturnStateEndRequest(t *testing.T) {
	// Mock the os.RemoveAll func
	defer Patch(&osRemoveAll, func(path string) error {
		return nil
	}).Restore()

	state := State{}

	newState, _, err := HandleStateCleanGit(state)

	if err != nil {
		t.Errorf("HandleStateCleanGit returned the error %q", err)
	}

	if newState != StateEndRequest {
		t.Errorf("HandleStateCleanGit returned %v instead of %v", newState, StateRunDocker)
	}
}

func TestHandleStateCleanGitShouldSetShouldCleanGitToFalse(t *testing.T) {
	// Mock the os.RemoveAll func
	defer Patch(&osRemoveAll, func(path string) error {
		return nil
	}).Restore()

	state := State{}

	_, newStatePayload, err := HandleStateCleanGit(state)

	if err != nil {
		t.Errorf("HandleStateCleanGit returned the error %q", err)
	}

	if newStatePayload.ShouldCleanGit {
		t.Errorf("HandleStateCleanGit should have set ShouldCleanGit to false")
	}
}
