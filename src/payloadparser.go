package main

import (
	"encoding/json"
	"errors"
	"strings"
)

const GITHUB_REFERENCE = "refs/heads/master"

type PushEvent struct {
	Ref string `json:"ref"`
}

func ParsePayload(payload []byte) (isValid bool, err error) {
	var pushEvent PushEvent

	err = json.Unmarshal(payload, &pushEvent)

	if err != nil {
		return false, err
	}

	if pushEvent.Ref == "" {
		return false, errors.New("invalid payload")
	}

	if !strings.EqualFold(pushEvent.Ref, GITHUB_REFERENCE) {
		return false, nil
	}

	return true, nil
}
