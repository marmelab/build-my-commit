package main

// TODO: restore this when Travis will supply a docker compatible environment
/* NOTE:
Currently, Travis run our tests inside a docker container running with an
old kernel not supported by docker.

Make sure to run those tests locally and to comment them before commiting
so that Travis pass.
*/

/*

import (
	"testing"
)

func TestDockerShouldFailWhenPassedInvalidCommand(t *testing.T) {
	err := docker("")

	if err == nil {
		t.Errorf("buildDocker() should have failed")
	}
}

func TestDockerShouldFailWhenPassedInvalidArguments(t *testing.T) {
	err := docker("build")

	if err == nil {
		t.Errorf("buildDocker() should have failed")
	}
}

func TestDockerShouldReturnWhenPassedValidBuildCommandAndArguments(t *testing.T) {
	err := docker("build", "-t", "test-build-docker", "test-build-docker")

	if err != nil {
		t.Errorf("buildDocker() should not have failed")
	}
}

func TestDockerShouldReturnWhenPassedValidRunCommandAndArguments(t *testing.T) {
	err := docker("run", "--name", "test-build-docker", "test-build-docker", "ls")

	if err != nil {
		t.Errorf("buildDocker() should not have failed")
	}
}
*/
