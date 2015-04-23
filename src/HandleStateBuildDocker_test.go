package main

import (
	"errors"
	"testing"
)

func TestHandleStateBuildDockerShouldReturnStateEndRequestWithErrorWhenDockerThrowsError(t *testing.T) {
	// Mock the exists func
	defer Patch(&docker, func(command string, arguments ...string) error {
		return errors.New("42")
	}).Restore()

	state := State{}

	newState, _, err := HandleStateBuildDocker(state)

	if newState != StateEndRequest {
		t.Errorf("HandleStateBuildDocker returned %v instead of %v", newState, StateEndRequest)
	}

	if err == nil {
		t.Error("HandleStateBuildDocker should have returned the error")
	}
}

func TestHandleStateBuildDockerShouldReturnStateRunDocker(t *testing.T) {
	// Mock the exists func
	defer Patch(&docker, func(command string, arguments ...string) error {
		return nil
	}).Restore()

	state := State{}

	newState, _, err := HandleStateBuildDocker(state)

	if err != nil {
		t.Errorf("HandleStateBuildDocker returned the error %q", err)
	}

	if newState != StateRunDocker {
		t.Errorf("HandleStateBuildDocker returned %v instead of %v", newState, StateRunDocker)
	}
}

func TestHandleStateBuildDockerShouldSetShouldCleanDockerToTrue(t *testing.T) {
	// Mock the exists func
	defer Patch(&docker, func(command string, arguments ...string) error {
		return nil
	}).Restore()

	state := State{}

	_, newStatePayload, err := HandleStateBuildDocker(state)

	if err != nil {
		t.Errorf("HandleStateBuildDocker returned the error %q", err)
	}

	if !newStatePayload.ShouldCleanDocker {
		t.Errorf("HandleStateBuildDocker should have set ShouldCleanDocker to true")
	}
}
