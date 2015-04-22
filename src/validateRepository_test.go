package main

import (
	"testing"
)

func TestValidateRepositoryShouldReturnFalseWhenRepositoryDoesNotContainDockerBuildFile(t *testing.T) {
	shouldProcess, err := validateRepository("tests/ValidateRepositoryNoDockerFile")

	if err != nil {
		t.Errorf("validateRepository() failed with error %q", err)
	}

	if shouldProcess {
		t.Errorf("validateRepository() should have returned false")
	}
}

func TestValidateRepositoryShouldReturnTrueWhenRepositoryContainDockerBuildFile(t *testing.T) {
	shouldProcess, err := validateRepository("tests/ValidateRepositoryWithDockerfile")

	if err != nil {
		t.Errorf("validateRepository() failed with error %q", err)
	}

	if !shouldProcess {
		t.Errorf("validateRepository() should have returned true")
	}
}
