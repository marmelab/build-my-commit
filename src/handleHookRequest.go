package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"regexp"
)

var gitCloneURLRegexp = regexp.MustCompile(`^https:\/\/github\.com\/.+\/(.+)\.git$`)

func writeResponse(responseWriter http.ResponseWriter, status int, msg string) {
	responseWriter.WriteHeader(status)
	responseWriter.Write([]byte(msg))
	return
}

func cleanRepository(path string) {
	os.RemoveAll(path)
}

func cleanDocker(containerName string) {
	docker("rm", containerName)
}

func handleHookRequest(responseWriter http.ResponseWriter, request *http.Request) {
	if request.Method != "POST" {
		writeResponse(responseWriter, http.StatusBadRequest, "Invalid method")
		return
	}

	body, err := ioutil.ReadAll(request.Body)

	if err != nil {
		writeResponse(responseWriter, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	pushEvent, err := parsePayload(body)

	if err != nil {
		writeResponse(responseWriter, http.StatusBadRequest, "Invalid payload")
		return
	}

	if pushEvent.Ref == "refs/heads/master" {
		// 1. Clone the project
		err := git(
			"clone",
			"--recursive",
			pushEvent.Repository.CloneURL)

		if err != nil {
			writeResponse(responseWriter, http.StatusInternalServerError, "Internal Server Error")
			cleanRepository(pushEvent.Repository.Name)
			return
		}

		// 2. Validate that we should process this repository
		shouldProcess, err := validateRepository(pushEvent.Repository.Name)

		if shouldProcess {
			// 3. Build the docker container
			err = docker(
				"build",
				fmt.Sprintf("--tag %q", pushEvent.Repository.Name),
				fmt.Sprintf("--file %q", path.Join(pushEvent.Repository.Name, dockerFilePath)),
				pushEvent.Repository.Name)

			if err != nil {
				cleanRepository(pushEvent.Repository.Name)
				cleanDocker(pushEvent.Repository.Name)
				writeResponse(responseWriter, http.StatusInternalServerError, "Internal Server Error")
				return
			}

			// 4. Run the docker container
			repositoryFullPath, err := filepath.Abs(pushEvent.Repository.Name)

			if err != nil {
				cleanRepository(pushEvent.Repository.Name)
				cleanDocker(pushEvent.Repository.Name)
				writeResponse(responseWriter, http.StatusInternalServerError, "Internal Server Error")
				return
			}

			err = docker(
				"run",
				fmt.Sprintf("--name %q", pushEvent.Repository.Name),
				fmt.Sprintf("--volume %q:/srv/", repositoryFullPath),
				pushEvent.Repository.Name,
				"make",
				"--file /srv/Makefile",
				"build")

			if err != nil {
				cleanRepository(pushEvent.Repository.Name)
				cleanDocker(pushEvent.Repository.Name)
				writeResponse(responseWriter, http.StatusInternalServerError, "Internal Server Error")
				return
			}

			// 5. Check wether the ouput has changed with git diff

			// 6. Commit & push the output if necessary
			cleanRepository(pushEvent.Repository.Name)
			cleanDocker(pushEvent.Repository.Name)
			writeResponse(responseWriter, http.StatusOK, "Repository processed")
		} else {
			cleanRepository(pushEvent.Repository.Name)
			cleanDocker(pushEvent.Repository.Name)
			writeResponse(responseWriter, http.StatusBadRequest, "Cannot process this repository")
			return
		}

		return
	}

	// Request should not be handled, just return
	return
}
