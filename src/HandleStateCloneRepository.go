package main

func HandleStateCloneRepository(state State) (int, State, error) {
	state.ShouldCleanGit = true

	_, err := git(
		"clone",
		"--recursive", // Ensure we pull all dependencies
		"--depth=1",   // Don't pull the whole history
		state.PushEvent.Repository.CloneURL,
		state.RepositoryPath)

	if err != nil {
		return StateEndRequest, state, err
	}

	return StateValidateRepository, state, nil
}
