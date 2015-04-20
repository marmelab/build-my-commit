package main

import (
	"testing"
)

func TestParsePayloadShouldReturnFalseAndErrorWhenNotSuppliedJson(t *testing.T) {
	payload := []byte(`invalid_json`)
	isValid, err := ParsePayload(payload)

	if err == nil {
		t.Errorf("ParsePayload() should have failed with error")
	}

	if isValid {
		t.Errorf("ParsePayload() should have returned false")
	}
}

func TestParsePayloadShouldReturnFalseAndErrorWhenSuppliedInvalidJson(t *testing.T) {
	payload := []byte(`{"toto": "invalid_json"}`)
	isValid, err := ParsePayload(payload)

	if err == nil {
		t.Errorf("ParsePayload() should have failed with error")
	}

	if isValid {
		t.Errorf("ParsePayload() should have returned false")
	}
}

func TestParsePayloadShouldReturnFalseAndErrorWhenSuppliedInvalidRefType(t *testing.T) {
	payload := []byte(`{"ref": 42}`)
	isValid, err := ParsePayload(payload)

	if err == nil {
		t.Errorf("ParsePayload() should have failed with error")
	}

	if isValid {
		t.Errorf("ParsePayload() should have returned false")
	}
}

func TestParsePayloadShouldReturnFalseWhenSuppliedValidGithubPayloadNotTargettingMaster(t *testing.T) {
	payload := []byte(`{"ref": "refs/heads/develop"}`)
	isValid, err := ParsePayload(payload)

	if err != nil {
		t.Errorf("ParsePayload() failed with error %q", err)
	}

	if isValid {
		t.Errorf("ParsePayload() should have returned false")
	}
}

func TestParsePayloadShouldReturnTrueWhenSuppliedValidGithubPayloadTargettingMaster(t *testing.T) {
	payload := []byte(`{"ref": "refs/heads/master"}`)
	isValid, err := ParsePayload(payload)

	if err != nil {
		t.Errorf("ParsePayload() failed with error %q", err)
	}

	if !isValid {
		t.Errorf("ParsePayload() should have returned false")
	}
}
