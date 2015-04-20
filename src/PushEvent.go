package main

type PushEvent struct {
	Ref string `json:"ref"`

    Repository struct {
        CloneUrl string `json:"clone_url"`
    } `json:"repository"`
}
