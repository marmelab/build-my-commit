package main

import (
	"io/ioutil"
	"net/http"
)

const GITHUB_REFERENCE = "refs/heads/master"

func returnError(responseWriter http.ResponseWriter, status int, msg string) {
	responseWriter.WriteHeader(status)
	responseWriter.Write([]byte(msg))
	return
}

func HandleHookRequest(responseWriter http.ResponseWriter, request *http.Request) {
	if request.Method != "POST" {
		returnError(responseWriter, http.StatusBadRequest, "Invalid method")
		return
	}

	body, err := ioutil.ReadAll(request.Body)

	if err != nil {
		returnError(responseWriter, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	pushEvent, err := ParsePayload(body)

	if err != nil {
		returnError(responseWriter, http.StatusBadRequest, "Invalid payload")
		return
	}

	if pushEvent.Ref == GITHUB_REFERENCE {
		// 1. Clone the project
		// TODO: name the returned path variable and use it in step 2
		_, err := CloneRepository(pushEvent.Repository.CloneUrl)

		if err != nil {
			returnError(responseWriter, http.StatusInternalServerError, "Internal Server Error")
			return
		}

		// 2. Validate that the project has a DockerFile

		// 3. Build the docker container

		// 4. Mount the docker container

		// 5. Check wether the ouput has changed with git diff

		// 6. Commit & push the output if necessary
		return
	} else {
		// Request should not be handled, just return
		return
	}
}
