package main

import (
	"net/http"
)

const StatusMessageForHandleStateRequestReceived = "Invalid HTTP method"

func HandleStateRequestReceived(state State) (int, State, error) {
	if state.Request.Method != "POST" {
		state.Status = http.StatusBadRequest
		state.StatusMessage = StatusMessageForHandleStateRequestReceived
		return StateEndRequest, state, nil
	}

	return StateParsePayload, state, nil
}
