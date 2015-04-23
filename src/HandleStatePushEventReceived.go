package main

import (
	"net/http"
)

const GitReferenceForMaster = "refs/heads/master"
const StatusMessageForStatePushEventReceivedNotOnMaster = "Discarded: not on master"

func HandleStatePushEventReceived(state State) (int, State, error) {
	if state.PushEvent.Ref == GitReferenceForMaster {
		return StateCloneRepository, state, nil
	}

	state.Status = http.StatusOK
	state.StatusMessage = StatusMessageForStatePushEventReceivedNotOnMaster
	return StateEndRequest, state, nil
}
