package main

import (
	"log"
	"os"
	"testing"
)

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

func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
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
