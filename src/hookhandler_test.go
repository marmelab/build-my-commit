package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHookHandlerShouldFailIfVerbIsNotPost(t *testing.T) {
	expected := 400
	request, err := http.NewRequest("GET", "http://example.com/foo", nil)

	if err != nil {
		t.Errorf("HookHandler() failed with error %q", err)
	}

	responseRecorder := httptest.NewRecorder()
	HookHandler(responseRecorder, request)

	status := responseRecorder.Code

	if status != expected {
		t.Errorf("HookHandler() should have failed with status %q for a GET request, returned %q", expected, status)
	}
}

func TestHookHandlerShouldFailWhenSentIncorrectJson(t *testing.T) {
	expected := 400
	body := []byte(`{"foo": "bar"}`)
	request, err := http.NewRequest("POST", "http://test", bytes.NewBuffer(body))
	request.Header.Set("Content-Type", "application/json")

	if err != nil {
		t.Errorf("HookHandler() failed with error %q", err)
	}

	responseRecorder := httptest.NewRecorder()
	HookHandler(responseRecorder, request)

	status := responseRecorder.Code

	if status != expected {
		t.Errorf("HookHandler() should have failed with status %q for a GET request, returned %q", expected, status)
	}
}
