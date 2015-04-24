package statehandlers

import (
	"github.com/marmelab/buildmycommit/states"
)

// RemoveDirectory is the signature of a function which removes a directory from the file system
type RemoveDirectory func(path string) error

// CleanRepository is the StateHandler in charge of the CleanRepository state
type CleanRepository struct {
	removeDirectory RemoveDirectory
}

// Handle the CleanRepository state
func (stateHandler CleanRepository) Handle(state states.State) (int, states.State) {
	err := stateHandler.removeDirectory(state.PushEvent.Repository.Name)

	if err != nil {
		return states.InternalServerError, state
	}

	state.ShouldCleanGit = false
	return states.EndRequest, state
}

// NewCleanRepository is a constructor like factory which returns a CleanRepository
func NewCleanRepository(removeDirectory RemoveDirectory) *CleanRepository {
	handler := new(CleanRepository)
	handler.removeDirectory = removeDirectory

	return handler
}
