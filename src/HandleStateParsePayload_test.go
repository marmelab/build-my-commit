package main

import (
	"errors"
	"io"
	"net/http"
	"testing"
)

func TestHandleStateParsePayloadShouldReturnStateEndRequestWithErrorWhenIOUtilThrowsError(t *testing.T) {
	// Mock the ioutil.ReadAll func
	defer Patch(&ioutilReadAll, func(r io.Reader) ([]byte, error) {
		return nil, errors.New("42")
	}).Restore()

	var request = http.Request{}
	state := State{Request: &request}

	newState, _, err := HandleStateParsePayload(state)

	if newState != StateEndRequest {
		t.Errorf("HandleStateParsePayload returned %v instead of %v", newState, StateEndRequest)
	}

	if err == nil {
		t.Error("HandleStateParsePayload should have returned the error")
	}
}

func TestHandleStateParsePayloadShouldReturnStateEndRequestWithErrorWhenParsePayloadThrowsError(t *testing.T) {
	// Mock the ioutil.ReadAll func
	defer Patch(&ioutilReadAll, func(r io.Reader) ([]byte, error) {
		return nil, nil
	}).Restore()

	// Mock the parsePayload func
	defer Patch(&parsePayload, func(payload []byte) (PushEvent, error) {
		return PushEvent{}, errors.New("42")
	}).Restore()

	var request = http.Request{}
	state := State{Request: &request}

	newState, _, err := HandleStateParsePayload(state)

	if newState != StateEndRequest {
		t.Errorf("HandleStateParsePayload returned %v instead of %v", newState, StateEndRequest)
	}

	if err == nil {
		t.Error("HandleStateParsePayload should have returned the error")
	}
}

func TestHandleStateParsePayloadShouldReturnStatePushEventReceived(t *testing.T) {
	// Mock the ioutil.ReadAll func
	defer Patch(&ioutilReadAll, func(r io.Reader) ([]byte, error) {
		return nil, nil
	}).Restore()

	// Mock the parsePayload func
	defer Patch(&parsePayload, func(payload []byte) (PushEvent, error) {
		return PushEvent{}, nil
	}).Restore()

	var request = http.Request{}
	state := State{Request: &request}

	newState, _, err := HandleStateParsePayload(state)

	if err != nil {
		t.Errorf("HandleStateParsePayload returned the error %q", err)
	}

	if newState != StatePushEventReceived {
		t.Errorf("HandleStateParsePayload returned %v instead of %v", newState, StatePushEventReceived)
	}
}
