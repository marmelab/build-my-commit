package main

import (
	"errors"
	"net/http"
	"testing"
)

func TestHandleStateVerifyRemoteHashShouldReturnStateEndRequestWithErrorWhenGitThrowsError(t *testing.T) {
	// Mock the gitWithContext func
	defer Patch(&gitWithContext, func(command string, contextPath string, arguments ...string) (string, error) {
		return "", errors.New("42")
	}).Restore()

	state := State{}

	newState, _, err := HandleStateVerifyRemoteHash(state)

	if newState != StateEndRequest {
		t.Errorf("HandleStateVerifyRemoteHash returned %v instead of %v", newState, StateEndRequest)
	}

	if err == nil {
		t.Error("HandleStateVerifyRemoteHash should have returned the error")
	}
}

func TestHandleStateVerifyRemoteHashShouldReturnStateEndRequestIfRemoteOutputHasChanged(t *testing.T) {
	// Mock the gitWithContext func
	defer Patch(&gitWithContext, func(command string, contextPath string, arguments ...string) (string, error) {
		return "has-changed", nil
	}).Restore()

	pushEvent := PushEvent{}
	pushEvent.HeadCommit.ID = "test"
	state := State{RepositoryPath: "repository-path", PushEvent: pushEvent}

	newState, _, err := HandleStateVerifyRemoteHash(state)

	if err != nil {
		t.Errorf("HandleStateVerifyRemoteHash returned the error %q", err)
	}

	if newState != StateEndRequest {
		t.Errorf("HandleStateVerifyRemoteHash returned %v instead of %v", newState, StateEndRequest)
	}
}

func TestHandleStateVerifyRemoteHashShouldSetStateStatusToOKRemoteOutputHasChanged(t *testing.T) {
	// Mock the gitWithContext func
	defer Patch(&gitWithContext, func(command string, contextPath string, arguments ...string) (string, error) {
		return "has-changed", nil
	}).Restore()

	pushEvent := PushEvent{}
	pushEvent.HeadCommit.ID = "test"
	state := State{RepositoryPath: "repository-path", PushEvent: pushEvent}

	_, newStatePayload, err := HandleStateVerifyRemoteHash(state)

	if err != nil {
		t.Errorf("HandleStateVerifyRemoteHash returned the error %q", err)
	}

	if newStatePayload.Status != http.StatusOK {
		t.Errorf("HandleStateVerifyRemoteHash returned %v instead of %v", newStatePayload.Status, http.StatusOK)
	}
}

func TestHandleStateVerifyRemoteHashShouldSetStateStatusMessageRemoteOutputHasChanged(t *testing.T) {
	// Mock the gitWithContext func
	defer Patch(&gitWithContext, func(command string, contextPath string, arguments ...string) (string, error) {
		return "has-changed", nil
	}).Restore()

	pushEvent := PushEvent{}
	pushEvent.HeadCommit.ID = "test"
	state := State{RepositoryPath: "repository-path", PushEvent: pushEvent}

	_, newStatePayload, err := HandleStateVerifyRemoteHash(state)

	if err != nil {
		t.Errorf("HandleStateVerifyRemoteHash returned the error %q", err)
	}

	if newStatePayload.StatusMessage != StatusMessageForHandleStateVerifyRemoteHash {
		t.Errorf("HandleStateVerifyRemoteHash returned %v instead of %v", newStatePayload.StatusMessage, StatusMessageForHandleStateVerifyRemoteHash)
	}
}

func TestHandleStateVerifyRemoteHashShouldReturnStateEndRequestIfRemoteOutputHasNotChanged(t *testing.T) {
	// Mock the gitWithContext func
	defer Patch(&gitWithContext, func(command string, contextPath string, arguments ...string) (string, error) {
		return "test", nil
	}).Restore()

	pushEvent := PushEvent{}
	pushEvent.HeadCommit.ID = "test"
	state := State{RepositoryPath: "repository-path", PushEvent: pushEvent}

	newState, _, err := HandleStateVerifyRemoteHash(state)

	if err != nil {
		t.Errorf("HandleStateVerifyRemoteHash returned the error %q", err)
	}

	if newState != StatePushOutput {
		t.Errorf("HandleStateVerifyRemoteHash returned %v instead of %v", newState, StatePushOutput)
	}
}
