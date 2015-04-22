package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleHookRequestShouldFailIfVerbIsNotPost(t *testing.T) {
	expected := 400
	request, err := http.NewRequest("GET", "http://example.com/foo", nil)

	if err != nil {
		t.Errorf("handleHookRequest() failed with error %q", err)
	}

	responseRecorder := httptest.NewRecorder()
	handleHookRequest(responseRecorder, request)

	status := responseRecorder.Code

	if status != expected {
		t.Errorf("handleHookRequest() should have failed with status %q for a GET request, returned %q", expected, status)
	}
}

func TestHandleHookRequestShouldFailWhenSentIncorrectJson(t *testing.T) {
	expected := 400
	body := []byte(`{"foo": "bar"}`)
	request, err := http.NewRequest("POST", "http://test", bytes.NewBuffer(body))
	request.Header.Set("Content-Type", "application/json")

	if err != nil {
		t.Errorf("handleHookRequest() failed with error %q", err)
	}

	responseRecorder := httptest.NewRecorder()
	handleHookRequest(responseRecorder, request)

	status := responseRecorder.Code

	if status != expected {
		t.Errorf("handleHookRequest() should have failed with status %q for a GET request, returned %q", expected, status)
	}
}
