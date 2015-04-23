package main

import (
	"testing"
)

func TestValidateRepositoryShouldReturnFalseWhenRepositoryDoesNotContainDockerBuildFile(t *testing.T) {
	// Mock the exists func
	defer Patch(&exists, func(path string) (bool, error) {
		return false, nil
	}).Restore()

	shouldProcess, err := validateRepository("tests/ValidateRepositoryNoDockerFile")

	if err != nil {
		t.Errorf("validateRepository() failed with error %q", err)
	}

	if shouldProcess {
		t.Errorf("validateRepository() should have returned false")
	}
}

func TestValidateRepositoryShouldReturnFalseWhenLastCommithasBeenMadeByTool(t *testing.T) {
	// Mock the exists func
	defer Patch(&exists, func(path string) (bool, error) {
		return true, nil
	}).Restore()

	// Mock the gitWithContext func
	defer Patch(&gitWithContext, func(command string, contextPath string, arguments ...string) (string, error) {
		return CommitMessage, nil
	}).Restore()

	shouldProcess, err := validateRepository("tests/ValidateRepositoryNoDockerFile")

	if err != nil {
		t.Errorf("validateRepository() failed with error %q", err)
	}

	if shouldProcess {
		t.Errorf("validateRepository() should have returned false")
	}
}

func TestValidateRepositoryShouldReturnTrueWhenRepositoryContainDockerBuildFile(t *testing.T) {
	// Mock the exists func
	defer Patch(&exists, func(path string) (bool, error) {
		return true, nil
	}).Restore()

	// Mock the gitWithContext func
	defer Patch(&gitWithContext, func(command string, contextPath string, arguments ...string) (string, error) {
		return "", nil
	}).Restore()

	shouldProcess, err := validateRepository("tests/ValidateRepositoryWithDockerfile")

	if err != nil {
		t.Errorf("validateRepository() failed with error %q", err)
	}

	if !shouldProcess {
		t.Errorf("validateRepository() should have returned true")
	}
}
