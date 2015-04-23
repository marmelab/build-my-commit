package main

import (
	"os"
)

var osRemoveAll = os.RemoveAll

func HandleStateCleanGit(state State) (int, State, error) {
	err := osRemoveAll(state.PushEvent.Repository.Name)

	if err != nil {
		return StateEndRequest, state, err
	}

	state.ShouldCleanGit = false
	return StateEndRequest, state, nil
}
