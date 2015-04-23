package main

import (
	"errors"
	"testing"
)

func TestHandleStateCommitOutputShouldReturnStateEndRequestWithErrorWhenGitThrowsError(t *testing.T) {
	// Mock the gitWithContext func
	defer Patch(&gitWithContext, func(command string, contextPath string, arguments ...string) (string, error) {
		return "", errors.New("42")
	}).Restore()

	state := State{}

	newState, _, err := HandleStateCommitOutput(state)

	if newState != StateEndRequest {
		t.Errorf("HandleStateCommitOutput returned %v instead of %v", newState, StateEndRequest)
	}

	if err == nil {
		t.Error("HandleStateCommitOutput should have returned the error")
	}
}

func TestHandleStateCommitOutputShouldReturnStateVerifyRemoteHash(t *testing.T) {
	// Mock the gitWithContext func
	defer Patch(&gitWithContext, func(command string, contextPath string, arguments ...string) (string, error) {
		return "", nil
	}).Restore()

	// Mock the getCommitMessage func
	defer Patch(&getCommitMessage, func(hash string) (string, error) {
		return "commit-message", nil
	}).Restore()

	state := State{RepositoryPath: "repository-path"}

	newState, _, err := HandleStateCommitOutput(state)

	if err != nil {
		t.Errorf("HandleStateCommitOutput returned the error %q", err)
	}

	if newState != StateVerifyRemoteHash {
		t.Errorf("HandleStateCommitOutput returned %v instead of %v", newState, StateVerifyRemoteHash)
	}
}
