package main

import (
	"testing"
)

func TestGenerateRandomIDReturnNewRandomAtEachCall(t *testing.T) {
	random1 := generateRandomID()
	random2 := generateRandomID()

	if random1 == random2 {
		t.Errorf("generateRandomID() should generate unique random strings")
	}
}
