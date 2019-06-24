package tools

import (
	"os"
	"testing"
)

func TestExistsReturnsTrueIfPathExists(t *testing.T) {
	path := os.TempDir()
	exists, _ := Exists(path)

	if !exists {
		t.Errorf("Exists should have return true for path %q", path)
	}
}

func TestExistsReturnsFalseIfPathDoesNotExists(t *testing.T) {
	path := "/42"
	exists, err := Exists(path)

	if exists {
		t.Errorf("Exists should have return false for path %q", path)
	}

	if err != nil {
		t.Errorf("Exists should not have returned an error for path %q", path)
	}
}
