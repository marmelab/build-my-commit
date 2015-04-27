package states

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
