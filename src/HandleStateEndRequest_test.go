package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleStateEndRequestShouldReturnStateCleanGitIfShouldCleanGitIsTrue(t *testing.T) {
	state := State{ShouldCleanGit: true, ResponseWriter: httptest.NewRecorder()}

	newState, _, err := HandleStateEndRequest(state)

	if err != nil {
		t.Errorf("HandleStateEndRequest returned the error %q", err)
	}

	if newState != StateCleanGit {
		t.Errorf("HandleStateEndRequest returned %v instead of %v", newState, StateCleanGit)
	}
}

func TestHandleStateEndRequestShouldReturnStateCleanDockerIfShouldCleanDockerIsTrue(t *testing.T) {
	state := State{ShouldCleanDocker: true, ResponseWriter: httptest.NewRecorder()}

	newState, _, err := HandleStateEndRequest(state)

	if err != nil {
		t.Errorf("HandleStateEndRequest returned the error %q", err)
	}

	if newState != StateCleanDocker {
		t.Errorf("HandleStateEndRequest returned %v instead of %v", newState, StateCleanDocker)
	}
}

func TestHandleStateEndRequestShouldReturnStateRequestHandledIfNothingToClean(t *testing.T) {
	state := State{ResponseWriter: httptest.NewRecorder()}

	newState, _, err := HandleStateEndRequest(state)

	if err != nil {
		t.Errorf("HandleStateEndRequest returned the error %q", err)
	}

	if newState != StateRequestHandled {
		t.Errorf("HandleStateEndRequest returned %v instead of %v", newState, StateRequestHandled)
	}
}

func TestHandleStateEndRequestShouldWriteStateStatus(t *testing.T) {
	recorder := httptest.NewRecorder()

	state := State{ResponseWriter: recorder, Status: http.StatusAccepted}

	_, _, err := HandleStateEndRequest(state)

	if err != nil {
		t.Errorf("HandleStateEndRequest returned the error %q", err)
	}

	if recorder.Code != http.StatusAccepted {
		t.Errorf("HandleStateEndRequest should have set status to %v but set it to %v", http.StatusAccepted, recorder.Code)
	}
}

func TestHandleStateEndRequestShouldWriteStateStatusMessage(t *testing.T) {
	recorder := httptest.NewRecorder()

	state := State{ResponseWriter: recorder, Status: http.StatusAccepted, StatusMessage: "foo"}

	_, _, err := HandleStateEndRequest(state)

	if err != nil {
		t.Errorf("HandleStateEndRequest returned the error %q", err)
	}

	if recorder.Body.String() != "foo" {
		t.Errorf("HandleStateEndRequest should have set status to %v but set it to %v", http.StatusAccepted, recorder.Code)
	}
}
