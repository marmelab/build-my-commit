package main

func HandleStatePushOutput(state State) (int, State, error) {
	_, err := gitWithContext(
		"push",
		state.RepositoryPath)

	if err != nil {
		return StateEndRequest, state, err
	}

	return StateEndRequest, state, nil
}
