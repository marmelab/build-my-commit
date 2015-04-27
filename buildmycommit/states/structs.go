package states

import (
	"net/http"
)

// State holds the data to pass between each states
type State struct {
	ResponseWriter    http.ResponseWriter
	Request           *http.Request
	PushEvent         PushEvent
	RepositoryPath    string
	Status            int
	StatusMessage     string
	ShouldCleanGit    bool
	ShouldCleanDocker bool
}
