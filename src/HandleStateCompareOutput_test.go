package main

import (
	"errors"
	"net/http"
	"testing"
)

func TestHandleStateCompareOutputShouldReturnStateEndRequestWithErrorWhenGitThrowsError(t *testing.T) {
	// Mock the gitWithContext func
	defer Patch(&gitWithContext, func(command string, contextPath string, arguments ...string) (string, error) {
		return "", errors.New("42")
	}).Restore()

	state := State{}

	newState, _, err := HandleStateCompareOutput(state)

	if newState != StateEndRequest {
		t.Errorf("HandleStateCompareOutput returned %v instead of %v", newState, StateEndRequest)
	}

	if err == nil {
		t.Error("HandleStateCompareOutput should have returned the error")
	}
}

func TestHandleStateCompareOutputShouldReturnStateEndRequestIfRemoteOutputIsUpToDate(t *testing.T) {
	// Mock the gitWithContext func
	defer Patch(&gitWithContext, func(command string, contextPath string, arguments ...string) (string, error) {
		// Return an empty string because we want no diff between local and remote
		return "", nil
	}).Restore()

	state := State{RepositoryPath: "repository-path"}

	newState, _, err := HandleStateCompareOutput(state)

	if err != nil {
		t.Errorf("HandleStateCompareOutput returned the error %q", err)
	}

	if newState != StateEndRequest {
		t.Errorf("HandleStateCompareOutput returned %v instead of %v", newState, StateEndRequest)
	}
}

func TestHandleStateCompareOutputShouldSetStateStatusToOKIfRemoteOutputIsUpToDate(t *testing.T) {
	// Mock the gitWithContext func
	defer Patch(&gitWithContext, func(command string, contextPath string, arguments ...string) (string, error) {
		// Return an empty string because we want no diff between local and remote
		return "", nil
	}).Restore()

	state := State{RepositoryPath: "repository-path"}

	_, newStatePayload, err := HandleStateCompareOutput(state)

	if err != nil {
		t.Errorf("HandleStateCompareOutput returned the error %q", err)
	}

	if newStatePayload.Status != http.StatusOK {
		t.Errorf("HandleStateCompareOutput returned %v instead of %v", newStatePayload.Status, http.StatusOK)
	}
}

func TestHandleStateCompareOutputShouldSetStateStatusMessageIfRemoteOutputIsUpToDate(t *testing.T) {
	// Mock the gitWithContext func
	defer Patch(&gitWithContext, func(command string, contextPath string, arguments ...string) (string, error) {
		// Return an empty string because we want no diff between local and remote
		return "", nil
	}).Restore()

	state := State{RepositoryPath: "repository-path"}

	_, newStatePayload, err := HandleStateCompareOutput(state)

	if err != nil {
		t.Errorf("HandleStateCompareOutput returned the error %q", err)
	}

	if newStatePayload.StatusMessage != StatusMessageForRemoteUpToDate {
		t.Errorf("HandleStateCompareOutput returned %v instead of %v", newStatePayload.StatusMessage, StatusMessageForRemoteUpToDate)
	}
}

func TestHandleStateCompareOutputShouldReturnStateEndRequestIfRemoteOutputIsNotUpToDate(t *testing.T) {
	// Mock the gitWithContext func
	defer Patch(&gitWithContext, func(command string, contextPath string, arguments ...string) (string, error) {
		// Return an non-empty string because we want some diff between local and remote
		return "status-indicating-diff-with-remote", nil
	}).Restore()

	state := State{RepositoryPath: "repository-path"}

	newState, _, err := HandleStateCompareOutput(state)

	if err != nil {
		t.Errorf("HandleStateCompareOutput returned the error %q", err)
	}

	if newState != StateCommitOutput {
		t.Errorf("HandleStateCompareOutput returned %v instead of %v", newState, StateCommitOutput)
	}
}
