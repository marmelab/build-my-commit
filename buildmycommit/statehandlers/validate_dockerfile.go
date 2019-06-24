package statehandlers

import (
	"github.com/marmelab/buildmycommit/states"
	"net/http"
	"path"
)

// FileExists is the signature of a function which check whether a file exists
type FileExists func(path string) (bool, error)

// ValidateDockerFile is the StateHandler in charge of the ValidateDockerFile state
type ValidateDockerFile struct {
	fileExists                    FileExists
	dockerFilePath                string
	statusMessageWhenNoDockerfile string
}

// Handle the ValidateDockerFile state
func (stateHandler ValidateDockerFile) Handle(state states.State) (int, states.State) {
	dockerFileFullPath := path.Join(state.RepositoryPath, stateHandler.dockerFilePath)

	dockerFileExists, err := stateHandler.fileExists(dockerFileFullPath)

	if err != nil {
		return states.InternalServerError, state
	}

	if dockerFileExists {
		return states.ValidateCommit, state
	}

	state.Status = http.StatusOK
	state.StatusMessage = stateHandler.statusMessageWhenNoDockerfile
	return states.EndRequest, state
}

// NewValidateDockerFile is a constructor like factory which returns a ValidateDockerFile
func NewValidateDockerFile(fileExists FileExists, dockerFilePath string, statusMessageWhenNoDockerfile string) *ValidateDockerFile {
	handler := new(ValidateDockerFile)
	handler.fileExists = fileExists
	handler.dockerFilePath = dockerFilePath
	handler.statusMessageWhenNoDockerfile = statusMessageWhenNoDockerfile

	return handler
}
