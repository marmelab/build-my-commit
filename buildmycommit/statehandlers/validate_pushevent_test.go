package statehandlers

import (
	"github.com/marmelab/buildmycommit/states"
	"net/http"
	"testing"
)

func TestPushEventReceivedShouldReturnStateCloneRepositoryWhenPushEventBranchMatchTheBranchToConsider(t *testing.T) {
	state := states.State{PushEvent: states.PushEvent{Ref: "branch-to-consider"}}

	handler := NewValidatePushEvent("branch-to-consider", "")

	newState, _ := handler.Handle(state)

	if newState != states.CloneRepository {
		t.Errorf("HandleStatePushEventReceived returned %v instead of %v", newState, states.CloneRepository)
	}
}

func TestPushEventReceivedShouldReturnStateEndRequestWhenPushEventBranchDoesNotMatchTheBranchToConsider(t *testing.T) {
	state := states.State{PushEvent: states.PushEvent{Ref: "not-the-good-branch"}}

	handler := NewValidatePushEvent("branch-to-consider", "")

	newState, _ := handler.Handle(state)

	if newState != states.EndRequest {
		t.Errorf("HandleStatePushEventReceived returned %v instead of %v", newState, states.EndRequest)
	}
}

func TestPushEventReceivedShouldSetStatusToOkWhenPushEventBranchDoesNotMatchTheBranchToConsider(t *testing.T) {
	state := states.State{PushEvent: states.PushEvent{Ref: "not-the-good-branch"}}

	handler := NewValidatePushEvent("branch-to-consider", "")

	_, newStatePayload := handler.Handle(state)

	if newStatePayload.Status != http.StatusOK {
		t.Errorf("HandleStatePushEventReceived returned %v instead of %v", newStatePayload.Status, http.StatusOK)
	}
}

func TestPushEventReceivedShouldSetStatusMessageWhenPushEventBranchDoesNotMatchTheBranchToConsider(t *testing.T) {
	state := states.State{PushEvent: states.PushEvent{Ref: "not-the-good-branch"}}

	handler := NewValidatePushEvent("branch-to-consider", "message")

	_, newStatePayload := handler.Handle(state)

	if newStatePayload.StatusMessage != "message" {
		t.Errorf("HandleStatePushEventReceived returned %v instead of %v", newStatePayload.StatusMessage, "message")
	}
}
