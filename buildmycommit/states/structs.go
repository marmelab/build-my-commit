package states

import (
	"net/http"
)

// PushEvent defines a struct mapping the github push event payload in json
type PushEvent struct {
	Ref string `json:"ref"`

	Repository struct {
		CloneURL string `json:"clone_url"`
		Name     string `json:"name"`
	} `json:"repository"`

	HeadCommit struct {
		ID string `json:"id"`
	} `json:"head_commit"`
}

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
