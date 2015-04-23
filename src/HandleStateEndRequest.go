package main

func HandleStateEndRequest(state State) (int, State, error) {
	state.ResponseWriter.WriteHeader(state.Status)
	state.ResponseWriter.Write([]byte(state.StatusMessage))

	if state.ShouldCleanGit {
		return StateCleanGit, state, nil
	}

	if state.ShouldCleanDocker {
		return StateCleanDocker, state, nil
	}

	return StateRequestHandled, state, nil
}
