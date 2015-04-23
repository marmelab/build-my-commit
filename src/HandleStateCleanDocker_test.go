package main

import (
	"errors"
	"testing"
)

func TestHandleStateCleanDockerShouldReturnStateEndRequestWithErrorWhenDockerThrowsError(t *testing.T) {
	// Mock the exists func
	defer Patch(&docker, func(command string, arguments ...string) error {
		return errors.New("42")
	}).Restore()

	state := State{}

	newState, _, err := HandleStateCleanDocker(state)

	if newState != StateEndRequest {
		t.Errorf("HandleStateCleanDocker returned %v instead of %v", newState, StateEndRequest)
	}

	if err == nil {
		t.Error("HandleStateCleanDocker should have returned the error")
	}
}

func TestHandleStateCleanDockerShouldReturnStateEndRequest(t *testing.T) {
	// Mock the exists func
	defer Patch(&docker, func(command string, arguments ...string) error {
		return nil
	}).Restore()

	state := State{}

	newState, _, err := HandleStateCleanDocker(state)

	if err != nil {
		t.Errorf("HandleStateCleanDocker returned the error %q", err)
	}

	if newState != StateEndRequest {
		t.Errorf("HandleStateCleanDocker returned %v instead of %v", newState, StateRunDocker)
	}
}

func TestHandleStateCleanDockerShouldSetShouldCleanDockerToFalse(t *testing.T) {
	// Mock the exists func
	defer Patch(&docker, func(command string, arguments ...string) error {
		return nil
	}).Restore()

	state := State{}

	_, newStatePayload, err := HandleStateCleanDocker(state)

	if err != nil {
		t.Errorf("HandleStateCleanDocker returned the error %q", err)
	}

	if newStatePayload.ShouldCleanDocker {
		t.Errorf("HandleStateCleanDocker should have set ShouldCleanDocker to false")
	}
}
