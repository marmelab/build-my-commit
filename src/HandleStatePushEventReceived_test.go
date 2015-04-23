package main

import (
	"net/http"
	"testing"
)

func TestHandleStatePushEventReceivedShouldReturnStateCloneRepositoryWhenPushEventBranchMatchMaster(t *testing.T) {
	state := State{PushEvent: PushEvent{Ref: GitReferenceForMaster}}

	newState, _, err := HandleStatePushEventReceived(state)

	if err != nil {
		t.Errorf("HandleStatePushEventReceived returned the error %q", err)
	}

	if newState != StateCloneRepository {
		t.Errorf("HandleStatePushEventReceived returned %v instead of %v", newState, StateCloneRepository)
	}
}

func TestHandleStatePushEventReceivedShouldReturnStateEndRequestWhenPushEventBranchDoesNotMatchMaster(t *testing.T) {
	state := State{PushEvent: PushEvent{Ref: "not-master"}}

	newState, _, err := HandleStatePushEventReceived(state)

	if err != nil {
		t.Errorf("HandleStatePushEventReceived returned the error %q", err)
	}

	if newState != StateEndRequest {
		t.Errorf("HandleStatePushEventReceived returned %v instead of %v", newState, StateEndRequest)
	}
}

func TestHandleStatePushEventReceivedShouldSetStatusToOkWhenPushEventBranchDoesNotMatchMaster(t *testing.T) {
	state := State{PushEvent: PushEvent{Ref: "not-master"}}

	_, newStatePayload, err := HandleStatePushEventReceived(state)

	if err != nil {
		t.Errorf("HandleStatePushEventReceived returned the error %q", err)
	}

	if newStatePayload.Status != http.StatusOK {
		t.Errorf("HandleStatePushEventReceived returned %v instead of %v", newStatePayload.Status, http.StatusOK)
	}
}

func TestHandleStatePushEventReceivedShouldSetStatusMessageWhenPushEventBranchDoesNotMatchMaster(t *testing.T) {
	state := State{PushEvent: PushEvent{Ref: "not-master"}}

	_, newStatePayload, err := HandleStatePushEventReceived(state)

	if err != nil {
		t.Errorf("HandleStatePushEventReceived returned the error %q", err)
	}

	if newStatePayload.StatusMessage != StatusMessageForStatePushEventReceivedNotOnMaster {
		t.Errorf("HandleStatePushEventReceived returned %v instead of %v", newStatePayload.StatusMessage, StatusMessageForStatePushEventReceivedNotOnMaster)
	}
}
