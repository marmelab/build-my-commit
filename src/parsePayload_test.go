package main

import (
	"testing"
)

func TestParsePayloadShouldFailWhenNotSuppliedJson(t *testing.T) {
	payload := []byte(`invalid_json`)
	pushEvent, err := parsePayload(payload)

	if err == nil {
		t.Errorf("parsePayload() should have failed with error")
	}

	if pushEvent.Ref != "" {
		t.Errorf("parsePayload() should have returned an empty pushEvent")
	}
}

func TestParsePayloadShouldFailWhenSuppliedInvalidJson(t *testing.T) {
	payload := []byte(`{"toto": "invalid_json"}`)
	pushEvent, err := parsePayload(payload)

	if err == nil {
		t.Errorf("parsePayload() should have failed with error")
	}

	if pushEvent.Ref != "" {
		t.Errorf("parsePayload() should have returned an empty pushEvent")
	}
}

func TestParsePayloadShouldFailWhenSuppliedInvalidRefType(t *testing.T) {
	payload := []byte(`{"ref": 42}`)
	pushEvent, err := parsePayload(payload)

	if err == nil {
		t.Errorf("parsePayload() should have failed with error")
	}

	if pushEvent.Ref != "" {
		t.Errorf("parsePayload() should have returned an empty pushEvent")
	}
}

func TestParsePayloadShouldReturnTrueWhenSuppliedValidGithubPayload(t *testing.T) {
	payload := []byte(`{"ref": "refs/heads/master"}`)
	pushEvent, err := parsePayload(payload)

	if err != nil {
		t.Errorf("parsePayload() failed with error %q", err)
	}

	if pushEvent.Ref != "refs/heads/master" {
		t.Errorf("parsePayload() should have returned a pushEvent with ref equal to %q", "refs/heads/master")
	}
}
