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
		_, err := git(
			"clone",
			"--recursive",
			pushEvent.Repository.CloneURL)

		if err != nil {
			writeResponse(responseWriter, http.StatusInternalServerError, "Internal Server Error")
			os.RemoveAll(pushEvent.Repository.Name)
			return
		}

		// 2. Validate that we should process this repository
		shouldProcess, err := validateRepository(pushEvent.Repository.Name)

		if shouldProcess {
			// 3. Fix permissions
			err = os.Chmod(pushEvent.Repository.Name, os.ModePerm)

			if err != nil {
				os.RemoveAll(pushEvent.Repository.Name)
				writeResponse(responseWriter, http.StatusInternalServerError, "Internal Server Error")
				return
			}

			// 4. Build the docker container
			err = docker(
				"build",
				fmt.Sprintf("--tag=%v", pushEvent.Repository.Name),
				fmt.Sprintf("--file=%v", path.Join(pushEvent.Repository.Name, dockerFilePath)),
				pushEvent.Repository.Name)

			if err != nil {
				os.RemoveAll(pushEvent.Repository.Name)
				writeResponse(responseWriter, http.StatusInternalServerError, "Internal Server Error")
				return
			}

			// 5. Run the docker container
			repositoryFullPath, err := filepath.Abs(pushEvent.Repository.Name)

			if err != nil {
				os.RemoveAll(pushEvent.Repository.Name)
				docker("rm", pushEvent.Repository.Name)
				writeResponse(responseWriter, http.StatusInternalServerError, "Internal Server Error")
				return
			}

			// docker run --name=test-repository-for-build-my-commit --volume=/home/gildas/projects/go/src/github.com/marmelab/build-my-commit/src/test-repository-for-build-my-commit/:/srv/ test-repository-for-build-my-commit make --file=/srv/Makefile
			err = docker(
				"run",
				fmt.Sprintf("--name=%v", pushEvent.Repository.Name),
				fmt.Sprintf("--volume=\"%v:/srv/\"", repositoryFullPath),
				pushEvent.Repository.Name,
				"make",
				"--file=/srv/Makefile",
				"build")

			if err != nil {
				os.RemoveAll(pushEvent.Repository.Name)
				docker("rm", pushEvent.Repository.Name)
				writeResponse(responseWriter, http.StatusInternalServerError, "Internal Server Error")
				return
			}

			// 6. Check wether the ouput has changed with git diff
			output, err := gitWithContext(
				"status",
				pushEvent.Repository.Name,
				"--porcelain")

			if err != nil {
				writeResponse(responseWriter, http.StatusInternalServerError, "Internal Server Error")
				os.RemoveAll(pushEvent.Repository.Name)
				docker("rm", pushEvent.Repository.Name)
				return
			}

			// 7. Commit & push the output if necessary
			if output != "" && len(output) > 0 {
				// Add files as it may be the first build
				_, err = gitWithContext(
					"add",
					pushEvent.Repository.Name,
					"build")

				if err != nil {
					writeResponse(responseWriter, http.StatusInternalServerError, "Internal Server Error")
					os.RemoveAll(pushEvent.Repository.Name)
					docker("rm", pushEvent.Repository.Name)
					return
				}

				// Commit files
				_, err = gitWithContext(
					"commit",
					pushEvent.Repository.Name,
					fmt.Sprintf("--message=%v", commitMessage),
					".")

				if err != nil {
					writeResponse(responseWriter, http.StatusInternalServerError, "Internal Server Error")
					os.RemoveAll(pushEvent.Repository.Name)
					docker("rm", pushEvent.Repository.Name)
					return
				}

				// Push files
				_, err = gitWithContext(
					"push",
					pushEvent.Repository.Name)

				if err != nil {
					writeResponse(responseWriter, http.StatusInternalServerError, "Internal Server Error")
					os.RemoveAll(pushEvent.Repository.Name)
					docker("rm", pushEvent.Repository.Name)
					return
				}
			}

			os.RemoveAll(pushEvent.Repository.Name)
			docker("rm", pushEvent.Repository.Name)
			writeResponse(responseWriter, http.StatusOK, "Repository processed")
		} else {
			os.RemoveAll(pushEvent.Repository.Name)
			writeResponse(responseWriter, http.StatusBadRequest, "Cannot process this repository")
			return
		}

		return
	}

	// Request should not be handled, just return
	return
}
