package main

import (
	"errors"
	"testing"
)

func TestHandleStateRunDockerShouldReturnStateEndRequestWithErrorWhenFilepathAbsThrowsError(t *testing.T) {
	// Mock the filepath.Abs func
	defer Patch(&filepathAbs, func(path string) (string, error) {
		return "", errors.New("42")
	}).Restore()

	state := State{}

	newState, _, err := HandleStateRunDocker(state)

	if newState != StateEndRequest {
		t.Errorf("HandleStateRunDocker returned %v instead of %v", newState, StateEndRequest)
	}

	if err == nil {
		t.Error("HandleStateRunDocker should have returned the error")
	}
}

func TestHandleStateRunDockerShouldReturnStateEndRequestWithErrorWhenDockerThrowsError(t *testing.T) {
	// Mock the filepath.Abs func
	defer Patch(&filepathAbs, func(path string) (string, error) {
		return "", nil
	}).Restore()

	// Mock the docker func
	defer Patch(&docker, func(command string, arguments ...string) error {
		return errors.New("42")
	}).Restore()

	state := State{}

	newState, _, err := HandleStateRunDocker(state)

	if newState != StateEndRequest {
		t.Errorf("HandleStateRunDocker returned %v instead of %v", newState, StateEndRequest)
	}

	if err == nil {
		t.Error("HandleStateRunDocker should have returned the error")
	}
}

func TestHandleStateRunDockerShouldReturnStateRunDocker(t *testing.T) {
	// Mock the filepath.Abs func
	defer Patch(&filepathAbs, func(path string) (string, error) {
		return "", nil
	}).Restore()

	// Mock the docker func
	defer Patch(&docker, func(command string, arguments ...string) error {
		return nil
	}).Restore()

	state := State{}

	newState, _, err := HandleStateRunDocker(state)

	if err != nil {
		t.Errorf("HandleStateRunDocker returned the error %q", err)
	}

	if newState != StateCompareOutput {
		t.Errorf("HandleStateRunDocker returned %v instead of %v", newState, StateCompareOutput)
	}
}
