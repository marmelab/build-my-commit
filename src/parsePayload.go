package main

import (
	"encoding/json"
	"errors"
)

var _parsePayload = func(payload []byte) (pushEvent PushEvent, err error) {
	err = json.Unmarshal(payload, &pushEvent)

	if err != nil {
		return pushEvent, err
	}

	if pushEvent.Ref == "" {
		return pushEvent, errors.New("invalid payload")
	}

	return pushEvent, nil
}

var parsePayload = _parsePayload
