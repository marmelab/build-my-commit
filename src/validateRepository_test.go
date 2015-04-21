package main

import (
	"os"
	"path"
	"testing"
)

func TestValidateRepositoryShouldReturnFalseWhenRepositoryDoesNotContainDockerBuildFile(t *testing.T) {
	path := "cloned-repository"
	err := os.Mkdir(path, os.ModeDir)

	if err != nil {
		t.Errorf("Could not create the test repository")
	}

	shouldProcess, err := validateRepository("cloned-repository")

	if err != nil {
		t.Errorf("validateRepository() failed with error %q", err)
	}

	if shouldProcess {
		t.Errorf("validateRepository() should have returned false")
	}

	os.RemoveAll(path)
}

func TestValidateRepositoryShouldReturnTrueWhenRepositoryContainDockerBuildFile(t *testing.T) {
	directoryPath := "cloned-repository"
	err := os.Mkdir(directoryPath, os.ModeDir)
	if err != nil {
		t.Errorf("Could not create the test repository")
	}

	_, err = os.Create(path.Join(directoryPath, "build.Dockerfile"))
	if err != nil {
		t.Errorf("Could not create the test repository")
	}

	shouldProcess, err := validateRepository("cloned-repository")

	if err != nil {
		t.Errorf("validateRepository() failed with error %q", err)
	}

	if !shouldProcess {
		t.Errorf("validateRepository() should have returned false")
	}

	os.RemoveAll(directoryPath)
}
