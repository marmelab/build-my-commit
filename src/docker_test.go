package main

import (
	"log"
	"os"
	"testing"
)

func TestCanCallDocker(t *testing.T) {
	if os.Getenv("CI") == "true" {
		log.Println("Skip docker test on Travis as it does not support running docker")
		return
	}

	err := docker("version")

	if err != nil {
		t.Errorf("docker('version') should not have failed")
	}
}
