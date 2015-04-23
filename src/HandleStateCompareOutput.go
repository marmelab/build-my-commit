package main

import (
	"net/http"
)

const StatusMessageForRemoteUpToDate = "Discarded: build already up to date on master"

func HandleStateCompareOutput(state State) (int, State, error) {
	output, err := gitWithContext(
		"status",
		state.RepositoryPath,
		"--porcelain") // This will make git status return a machine readable output without pretty formatting

	if err != nil {
		return StateEndRequest, state, err
	}

	// 7. Commit & push the output if necessary
	if output != "" && len(output) > 0 {
		return StateCommitOutput, state, nil
	}

	state.Status = http.StatusOK
	state.StatusMessage = StatusMessageForRemoteUpToDate
	return StateEndRequest, state, nil
}
