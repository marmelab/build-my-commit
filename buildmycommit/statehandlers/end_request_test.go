package statehandlers

import (
	"github.com/marmelab/buildmycommit/states"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestEndRequestShouldReturnStateCleanGitIfShouldCleanGitIsTrue(t *testing.T) {
	state := states.State{ShouldCleanGit: true, ResponseWriter: httptest.NewRecorder()}

	handler := NewEndRequest()

	newState, _ := handler.Handle(state)

	if newState != states.CleanRepository {
		t.Errorf("EndRequest returned %v instead of %v", newState, states.CleanRepository)
	}
}

func TestEndRequestShouldReturnStateCleanDockerIfShouldCleanDockerIsTrue(t *testing.T) {
	state := states.State{ShouldCleanDocker: true, ResponseWriter: httptest.NewRecorder()}

	handler := NewEndRequest()

	newState, _ := handler.Handle(state)

	if newState != states.CleanDocker {
		t.Errorf("EndRequest returned %v instead of %v", newState, states.CleanDocker)
	}
}

func TestEndRequestShouldReturnStateRequestHandledIfNothingToClean(t *testing.T) {
	state := states.State{ResponseWriter: httptest.NewRecorder()}

	handler := NewEndRequest()

	newState, _ := handler.Handle(state)

	if newState != states.RequestHandled {
		t.Errorf("EndRequest returned %v instead of %v", newState, states.RequestHandled)
	}
}

func TestEndRequestShouldWriteStateStatus(t *testing.T) {
	recorder := httptest.NewRecorder()

	state := states.State{ResponseWriter: recorder, Status: http.StatusAccepted}

	handler := NewEndRequest()

	handler.Handle(state)

	if recorder.Code != http.StatusAccepted {
		t.Errorf("EndRequest should have set status to %v but set it to %v", http.StatusAccepted, recorder.Code)
	}
}

func TestEndRequestShouldWriteStateStatusMessage(t *testing.T) {
	recorder := httptest.NewRecorder()

	state := states.State{ResponseWriter: recorder, Status: http.StatusAccepted, StatusMessage: "foo"}

	handler := NewEndRequest()

	handler.Handle(state)

	if recorder.Body.String() != "foo" {
		t.Errorf("EndRequest should have set status to %v but set it to %v", http.StatusAccepted, recorder.Code)
	}
}
