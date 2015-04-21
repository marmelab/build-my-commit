package main

type pushEvent struct {
	Ref string `json:"ref"`

	Repository struct {
		CloneURL string `json:"clone_url"`
		Name     string `json:"name"`
	} `json:"repository"`
}
