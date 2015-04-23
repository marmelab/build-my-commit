package main

func HandleStateCleanDocker(state State) (int, State, error) {
	err := docker("rm", state.PushEvent.Repository.Name)

	if err != nil {
		return StateEndRequest, state, err
	}

	state.ShouldCleanDocker = false
	return StateEndRequest, state, nil
}
