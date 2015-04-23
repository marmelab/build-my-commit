package main

import (
	"fmt"
)

func HandleStateCommitOutput(state State) (int, State, error) {
	// Add files as it may be the first build
	_, err := gitWithContext(
		"add",
		state.RepositoryPath,
		"build")

	if err != nil {
		return StateEndRequest, state, err
	}

	// Commit files
	message, err := getCommitMessage(state.PushEvent.HeadCommit.ID)

	if err != nil {
		return StateEndRequest, state, err
	}

	_, err = gitWithContext(
		"commit",
		state.RepositoryPath,
		fmt.Sprintf("--message=%v", message),
		".")

	if err != nil {
		return StateEndRequest, state, err
	}

	return StateVerifyRemoteHash, state, nil
}
