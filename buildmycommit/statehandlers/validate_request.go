package statehandlers

import (
	"github.com/marmelab/buildmycommit/states"
	"net/http"
)

// ValidateRequest is the StateHandler in charge of the ValidateRequest state
type ValidateRequest struct {
	statusMessageWhenInvalidHTTPMethod string
}

// Handle the ValidateRequest state
func (stateHandler ValidateRequest) Handle(state states.State) (int, states.State) {
	if state.Request.Method != "POST" {
		state.Status = http.StatusBadRequest
		state.StatusMessage = stateHandler.statusMessageWhenInvalidHTTPMethod
		return states.EndRequest, state
	}

	return states.ParsePayload, state
}

// NewValidateRequest is a constructor like factory which returns a ValidateRequest
func NewValidateRequest(statusMessageWhenInvalidHTTPMethod string) *ValidateRequest {
	handler := new(ValidateRequest)
	handler.statusMessageWhenInvalidHTTPMethod = statusMessageWhenInvalidHTTPMethod

	return handler
}
