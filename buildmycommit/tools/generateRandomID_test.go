package tools

import (
	"testing"
)

func TestGenerateRandomIDReturnNewRandomAtEachCall(t *testing.T) {
	random1 := GenerateRandomID()
	random2 := GenerateRandomID()

	if random1 == random2 {
		t.Errorf("generateRandomID() should generate unique random strings")
	}
}
