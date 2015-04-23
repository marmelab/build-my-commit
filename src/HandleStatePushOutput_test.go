package main

import (
	"errors"
	"testing"
)

func TestHandleStatePushOutputShouldReturnStateEndRequestWithErrorWhenGitThrowsError(t *testing.T) {
	// Mock the gitWithContext func
	defer Patch(&gitWithContext, func(command string, contextPath string, arguments ...string) (string, error) {
		return "", errors.New("42")
	}).Restore()

	state := State{}

	newState, _, err := HandleStatePushOutput(state)

	if newState != StateEndRequest {
		t.Errorf("HandleStatePushOutput returned %v instead of %v", newState, StateEndRequest)
	}

	if err == nil {
		t.Error("HandleStatePushOutput should have returned the error")
	}
}

func TestHandleStatePushOutputShouldReturnStateVerifyRemoteHash(t *testing.T) {
	// Mock the gitWithContext func
	defer Patch(&gitWithContext, func(command string, contextPath string, arguments ...string) (string, error) {
		return "", nil
	}).Restore()

	// Mock the getCommitMessage func
	defer Patch(&getCommitMessage, func(hash string) (string, error) {
		return "", nil
	}).Restore()

	state := State{RepositoryPath: "repository-path"}

	newState, _, err := HandleStatePushOutput(state)

	if err != nil {
		t.Errorf("HandleStatePushOutput returned the error %q", err)
	}

	if newState != StateEndRequest {
		t.Errorf("HandleStatePushOutput returned %v instead of %v", newState, StateEndRequest)
	}
}
