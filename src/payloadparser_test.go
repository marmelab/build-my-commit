package main

import (
	"testing"
)

func TestParsePayloadShouldFailWhenNotSuppliedJson(t *testing.T) {
	payload := []byte(`invalid_json`)
	pushEvent, err := ParsePayload(payload)

	if err == nil {
		t.Errorf("ParsePayload() should have failed with error")
	}

	if pushEvent.Ref != "" {
		t.Errorf("ParsePayload() should have returned an empty PushEvent")
	}
}

func TestParsePayloadShouldFailWhenSuppliedInvalidJson(t *testing.T) {
	payload := []byte(`{"toto": "invalid_json"}`)
	pushEvent, err := ParsePayload(payload)

	if err == nil {
		t.Errorf("ParsePayload() should have failed with error")
	}

	if pushEvent.Ref != "" {
		t.Errorf("ParsePayload() should have returned an empty PushEvent")
	}
}

func TestParsePayloadShouldFailWhenSuppliedInvalidRefType(t *testing.T) {
	payload := []byte(`{"ref": 42}`)
	pushEvent, err := ParsePayload(payload)

	if err == nil {
		t.Errorf("ParsePayload() should have failed with error")
	}

	if pushEvent.Ref != "" {
		t.Errorf("ParsePayload() should have returned an empty PushEvent")
	}
}

func TestParsePayloadShouldReturnTrueWhenSuppliedValidGithubPayload(t *testing.T) {
	payload := []byte(`{"ref": "refs/heads/master"}`)
	pushEvent, err := ParsePayload(payload)

	if err != nil {
		t.Errorf("ParsePayload() failed with error %q", err)
	}

	if pushEvent.Ref != "refs/heads/master" {
		t.Errorf("ParsePayload() should have returned a PushEvent with ref equal to %q", "refs/heads/master")
	}
}
