package statehandlers

import (
	"github.com/marmelab/buildmycommit/states"
	"net/http"
)

// InternalServerError is the StateHandler in charge of the InternalServerError state
type InternalServerError struct {
}

// Handle the InternalServerError state
func (stateHandler InternalServerError) Handle(state states.State) (int, states.State) {
	state.Status = http.StatusInternalServerError
	state.StatusMessage = http.StatusText(state.Status)
	return states.EndRequest, state
}

// NewInternalServerError is a constructor like factory which returns a InternalServerError
func NewInternalServerError() *InternalServerError {
	handler := new(InternalServerError)

	return handler
}
