package statehandlers

import (
	"github.com/marmelab/buildmycommit/states"
	"net/http"
	"testing"
)

func TestRequestReceivedShouldReturnStateEndRequestWhenMethodIsNotPOST(t *testing.T) {
	var request = http.Request{Method: "GET"}
	state := states.State{Request: &request}

	handler := NewValidateRequest("")

	newState, _ := handler.Handle(state)

	if newState != states.EndRequest {
		t.Errorf("HandleStateRequestReceived returned %v instead of %v", newState, states.EndRequest)
	}
}

func TestRequestReceivedShouldSetStatusToStatusBadRequestWhenMethodIsNotPOST(t *testing.T) {
	var request = http.Request{Method: "GET"}
	state := states.State{Request: &request}

	handler := NewValidateRequest("")

	_, newStatePayload := handler.Handle(state)

	if newStatePayload.Status != http.StatusBadRequest {
		t.Errorf("HandleStateRequestReceived returned %v instead of %v", newStatePayload.Status, http.StatusBadRequest)
	}
}

func TestRequestReceivedShouldSetStatusMessageWhenMethodIsNotPOST(t *testing.T) {
	var request = http.Request{Method: "GET"}
	state := states.State{Request: &request}

	handler := NewValidateRequest("message")

	_, newStatePayload := handler.Handle(state)

	if newStatePayload.StatusMessage != "message" {
		t.Errorf("HandleStateRequestReceived returned %v instead of %v", newStatePayload.StatusMessage, "message")
	}
}

func TestRequestReceivedShouldReturnStateParsePayload(t *testing.T) {
	var request = http.Request{Method: "POST"}
	state := states.State{Request: &request}

	handler := NewValidateRequest("")

	newState, _ := handler.Handle(state)

	if newState != states.ParsePayload {
		t.Errorf("HandleStateRequestReceived returned %v instead of %v", newState, states.ParsePayload)
	}
}
