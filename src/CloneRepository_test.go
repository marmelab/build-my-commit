package main

import (
	"testing"
	"os"
)

type TestRunner struct{}

func TestCloneRepositoryShouldFailWhenPassedEmptyUrl(t *testing.T) {
    _, err := CloneRepository("")

    if err == nil {
        t.Errorf("CloneRepository() should have failed")
    }
}

func TestCloneRepositoryShouldFailWhenPassedInvalidGitUrl(t *testing.T) {
	_, err := CloneRepository("http://google.com")

    if err == nil {
        t.Errorf("CloneRepository() should have failed")
    }
}

func TestCloneRepositoryShouldPassWhenPassedValidGitUrl(t *testing.T) {
    path, err := CloneRepository("https://github.com/djhi/test-repository-for-build-my-commit.git")

    if err != nil {
        t.Errorf("CloneRepository() should have failed: %q", err)
    }

	directoryExist, err := exists(path)
	if !directoryExist {
        t.Errorf("CloneRepository() should have cloned the repository: %q", err)
	}

	err = os.RemoveAll(path)
	if err != nil {
		t.Errorf("Unable to remove cloned repository: %q", err)
	}

}
