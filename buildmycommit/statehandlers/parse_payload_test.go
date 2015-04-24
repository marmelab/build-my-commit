package statehandlers

import (
	"errors"
	"github.com/marmelab/buildmycommit/states"
	"io"
	"net/http"
	"testing"
)

func TestParsePayloadShouldReturnStateInternalServerErrorWhenIOUtilThrowsError(t *testing.T) {
	// Mock the readPayloadFromBuffer func
	readPayloadFromBuffer := func(r io.Reader) ([]byte, error) {
		return nil, errors.New("42")
	}

	var request = http.Request{}
	state := states.State{Request: &request}

	handler := NewParsePayload(readPayloadFromBuffer, nil, "")

	newState, _ := handler.Handle(state)

	if newState != states.InternalServerError {
		t.Errorf("HandleStateParsePayload returned %v instead of %v", newState, states.InternalServerError)
	}
}

func TestParsePayloadShouldReturnStateInternalServerErrorWhenJsonUnmarshalThrowsError(t *testing.T) {
	// Mock the readPayloadFromBuffer func
	readPayloadFromBuffer := func(r io.Reader) ([]byte, error) {
		return nil, nil
	}

	// Mock the unmarshalJSON func
	unmarshalJSON := func(data []byte, v interface{}) error {
		return errors.New("42")
	}

	var request = http.Request{}
	state := states.State{Request: &request}

	handler := NewParsePayload(readPayloadFromBuffer, unmarshalJSON, "")

	newState, _ := handler.Handle(state)

	if newState != states.InternalServerError {
		t.Errorf("HandleStateParsePayload returned %v instead of %v", newState, states.InternalServerError)
	}
}

func TestParsePayloadShouldReturnStatePushEventReceived(t *testing.T) {
	// Mock the readPayloadFromBuffer func
	readPayloadFromBuffer := func(r io.Reader) ([]byte, error) {
		return nil, nil
	}

	// Mock the unmarshalJSON func
	unmarshalJSON := func(data []byte, v interface{}) error {
		return nil
	}

	var request = http.Request{}
	state := states.State{Request: &request}

	handler := NewParsePayload(readPayloadFromBuffer, unmarshalJSON, "")

	newState, _ := handler.Handle(state)

	if newState != states.ValidatePushEvent {
		t.Errorf("HandleStateParsePayload returned %v instead of %v", newState, states.ValidatePushEvent)
	}
}
