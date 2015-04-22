package main

import (
	"testing"
)

func TestCanCallGit(t *testing.T) {
	_, err := git("version", "")

	if err != nil {
		t.Errorf("git('version') should failed with error %q", err)
	}
}
