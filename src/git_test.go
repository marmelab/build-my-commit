package main

import (
	"testing"
)

func TestCanCallGit(t *testing.T) {
	err := git("version", "")

	if err != nil {
		t.Errorf("git('version') should not have failed")
	}
}
