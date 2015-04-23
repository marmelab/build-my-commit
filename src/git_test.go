package main

import (
	"testing"
)

func TestCanCallGit(t *testing.T) {
	_, err := git("version", "")

	if err != nil {
		t.Errorf("git('version') have failed with error %q", err)
	}
}

func TestGetCommitMessageReturnEmptyStringAndErrorWhenPassedEmptyHash(t *testing.T) {
	msg, err := getCommitMessage("")

	if err == nil {
		t.Errorf("getCommitMessage('') should failed with error")
	}

	if msg != "" {
		t.Errorf("getCommitMessage('') should have returned an empty string")
	}
}

func TestGetCommitMessageReturnDefaultCommitMessageAndHashWhenPassedValidHash(t *testing.T) {
	hash := "hash"
	msg, err := getCommitMessage(hash)

	if err != nil {
		t.Errorf("getCommitMessage('hash') have failed with error %q", err)
	}

	if msg != commitMessage+" "+hash {
		t.Errorf("getCommitMessage('hash') should have returned \"%q\" but returned %q", commitMessage+" "+hash, msg)
	}
}
