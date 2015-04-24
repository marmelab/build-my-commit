package statehandlers

import (
	"github.com/marmelab/buildmycommit/states"
	"net/http"
	"testing"
)

func TestInternalServerErrorShouldSetStatusToStatusInternalServerError(t *testing.T) {
	state := states.State{}

	handler := NewInternalServerError()

	_, payload := handler.Handle(state)

	if payload.Status != http.StatusInternalServerError {
		t.Errorf("InternalServerError should have set status to %v but set it to %v", http.StatusInternalServerError, payload.Status)
	}
}

func TestInternalServerErrorShouldWriteSetStatusMessageToInternalServerError(t *testing.T) {
	state := states.State{}

	handler := NewInternalServerError()

	_, payload := handler.Handle(state)

	expectedMessage := http.StatusText(http.StatusInternalServerError)
	if payload.StatusMessage != expectedMessage {
		t.Errorf("InternalServerError should have set status to %v but set it to %v", expectedMessage, payload.StatusMessage)
	}
}
