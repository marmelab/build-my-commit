package main

import (
	"log"
	"net/http"
)

const StatusMessageForHandleStateVerifyRemoteHash = "Discarded: master has been updated"

func HandleStateVerifyRemoteHash(state State) (int, State, error) {
	lastHash, err := gitWithContext(
		"rev-parse",
		state.RepositoryPath,
		"origin/master")

	if err != nil {
		return StateEndRequest, state, err
	}

	log.Printf("Last commit hash: \"%q\"", lastHash)
	log.Printf("Pushevent commit hash: \"%q\"", state.PushEvent.HeadCommit.ID)

	// Repository hasn't been updated if we get the same hash as pushEvent
	// If not, we just discard this build as another one is probably running in parallel
	if lastHash == state.PushEvent.HeadCommit.ID {
		return StatePushOutput, state, nil
	}

	state.Status = http.StatusOK
	state.StatusMessage = StatusMessageForHandleStateVerifyRemoteHash
	return StateEndRequest, state, nil
}
