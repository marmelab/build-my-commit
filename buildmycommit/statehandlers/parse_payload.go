package statehandlers

import (
	"github.com/marmelab/buildmycommit/states"
	"io"
	"net/http"
)

// ReadPayloadFromBuffer is the signature of a function which reads http payload from an io.Reader
type ReadPayloadFromBuffer func(r io.Reader) ([]byte, error)

// UnmarshalJSON is the signature of a function which parses unmarshal an object from a json payload
type UnmarshalJSON func(b []byte, v interface{}) error

// ParsePayload is the StateHandler in charge of the ParsePayload state
type ParsePayload struct {
	readPayloadFromBuffer             ReadPayloadFromBuffer
	unmarshalJSON                     UnmarshalJSON
	statusMessageWhenPayloadIsInvalid string
}

// Handle the ParsePayload state
func (stateHandler ParsePayload) Handle(state states.State) (int, states.State) {
	payload, err := stateHandler.readPayloadFromBuffer(state.Request.Body)

	if err != nil {
		return states.InternalServerError, state
	}

	pushEvent := states.PushEvent{}

	err = stateHandler.unmarshalJSON(payload, &pushEvent)

	if err != nil {
		state.Status = http.StatusBadRequest
		state.StatusMessage = stateHandler.statusMessageWhenPayloadIsInvalid
		return states.InternalServerError, state
	}

	state.PushEvent = pushEvent
	return states.ValidatePushEvent, state
}

// NewParsePayload is a constructor like factory which returns a ParsePayload
func NewParsePayload(readPayloadFromBuffer ReadPayloadFromBuffer, unmarshalJSON UnmarshalJSON, statusMessageWhenPayloadIsInvalid string) *ParsePayload {
	handler := new(ParsePayload)
	handler.readPayloadFromBuffer = readPayloadFromBuffer
	handler.unmarshalJSON = unmarshalJSON
	handler.statusMessageWhenPayloadIsInvalid = statusMessageWhenPayloadIsInvalid
	return handler
}
