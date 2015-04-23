package main

import (
	"net/http"
	"testing"
)

func TestHandleStateRequestReceivedShouldReturnStateEndRequestWhenMethodIsNotPOST(t *testing.T) {
	var request = http.Request{Method: "GET"}
	state := State{Request: &request}

	newState, _, err := HandleStateRequestReceived(state)

	if err != nil {
		t.Errorf("HandleStateRequestReceived returned the error %q", err)
	}

	if newState != StateEndRequest {
		t.Errorf("HandleStateRequestReceived returned %v instead of %v", newState, StateEndRequest)
	}
}

func TestHandleStateRequestReceivedShouldSetStatusToStatusBadRequestWhenMethodIsNotPOST(t *testing.T) {
	var request = http.Request{Method: "GET"}
	state := State{Request: &request}

	_, newStatePayload, err := HandleStateRequestReceived(state)

	if err != nil {
		t.Errorf("HandleStateRequestReceived returned the error %q", err)
	}

	if newStatePayload.Status != http.StatusBadRequest {
		t.Errorf("HandleStateRequestReceived returned %v instead of %v", newStatePayload.Status, http.StatusBadRequest)
	}
}

func TestHandleStateRequestReceivedShouldSetStatusMessageWhenMethodIsNotPOST(t *testing.T) {
	var request = http.Request{Method: "GET"}
	state := State{Request: &request}

	_, newStatePayload, err := HandleStateRequestReceived(state)

	if err != nil {
		t.Errorf("HandleStateRequestReceived returned the error %q", err)
	}

	if newStatePayload.StatusMessage != StatusMessageForHandleStateRequestReceived {
		t.Errorf("HandleStateRequestReceived returned %v instead of %v", newStatePayload.StatusMessage, StatusMessageForHandleStateRequestReceived)
	}
}

func TestHandleStateRequestReceivedShouldReturnStateParsePayload(t *testing.T) {
	var request = http.Request{Method: "POST"}
	state := State{Request: &request}

	newState, _, err := HandleStateRequestReceived(state)

	if err != nil {
		t.Errorf("HandleStateRequestReceived returned the error %q", err)
	}

	if newState != StateParsePayload {
		t.Errorf("HandleStateRequestReceived returned %v instead of %v", newState, StateParsePayload)
	}
}
