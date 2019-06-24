package statehandlers

import (
	"github.com/marmelab/buildmycommit/states"
	"github.com/marmelab/buildmycommit/git"
)

// CloneRepository is the StateHandler in charge of the CloneRepository state
type CloneRepository struct {
	executeGit git.ExecuteGit
}

// Handle the CloneRepository state
func (stateHandler CloneRepository) Handle(state states.State) (int, states.State) {
	state.ShouldCleanGit = true

	_, err := stateHandler.executeGit(
		"clone",
		"--recursive", // Ensure we pull all dependencies
		"--depth=1",   // Don't pull the whole history
		state.PushEvent.Repository.CloneURL,
		state.RepositoryPath)

	if err != nil {
		return states.InternalServerError, state
	}

	return states.ValidateDockerFile, state
}

// NewCloneRepository is a constructor like factory which returns a CloneRepository
func NewCloneRepository(executeGit git.ExecuteGit) *CloneRepository {
	handler := new(CloneRepository)
	handler.executeGit = executeGit

	return handler
}
