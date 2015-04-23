package main

import (
	"io/ioutil"
)

var ioutilReadAll = ioutil.ReadAll

func HandleStateParsePayload(state State) (int, State, error) {
	body, err := ioutilReadAll(state.Request.Body)

	if err != nil {
		return StateEndRequest, state, err
	}

	pushEvent, err := parsePayload(body)

	if err != nil {
		return StateEndRequest, state, err
	}

	state.PushEvent = pushEvent
	return StatePushEventReceived, state, nil
}
