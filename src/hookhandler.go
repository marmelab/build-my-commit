package main

import (
	"io/ioutil"
	"net/http"
)

func returnError(responseWriter http.ResponseWriter, status int, msg string) {
	responseWriter.WriteHeader(http.StatusBadRequest)
	responseWriter.Write([]byte(msg))
	return
}

func HookHandler(responseWriter http.ResponseWriter, request *http.Request) {
	if request.Method != "POST" {
		returnError(responseWriter, 400, "Invalid method")
		return
	}

	body, err := ioutil.ReadAll(request.Body)

	if err != nil {
		returnError(responseWriter, 500, "")
		return
	}

	isValid, err := ParsePayload(body)

	if err != nil {
		returnError(responseWriter, 400, "Invalid method")
		return
	}

	if isValid {
		// Request should be handled

		// 1. Clone the project

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
