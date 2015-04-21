package main

import (
	"os"
	"testing"
)

func TestGitShouldFailWhenPassedInvalidCommand(t *testing.T) {
	err := git("")

	if err == nil {
		t.Errorf("git() should have failed")
	}
}

func TestGitShouldFailWhenPassedInvalidArguments(t *testing.T) {
	err := git("clone", "")

	if err == nil {
		t.Errorf("git() should have failed")
	}
}

func TestGitShouldPassWhenPassedValidGitUrl(t *testing.T) {
	path := "test-repository-for-build-my-commit"
	err := git("clone", "https://github.com/djhi/test-repository-for-build-my-commit.git")

	if err != nil {
		t.Errorf("git() should have failed: %q", err)
	}

	directoryExist, err := exists(path)
	if !directoryExist {
		t.Errorf("git() should have cloned the repository: %q", err)
	}

	err = os.RemoveAll(path)
	if err != nil {
		t.Errorf("Unable to remove cloned repository: %q", err)
	}

}
