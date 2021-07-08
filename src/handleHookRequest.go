package main

import (
	"fmt"
	"io/ioutil"
	"log"
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
		repositoryPath := generateRandomID()

		// 1. Clone the project
		_, err := git(
			"clone",
			"--recursive", // Ensure we pull all dependencies
			"--depth=1",   // Don't pull the whole history
			pushEvent.Repository.CloneURL,
			repositoryPath)

		if err != nil {
			writeResponse(responseWriter, http.StatusInternalServerError, "Internal Server Error")
			os.RemoveAll(repositoryPath)
			return
		}

		// 2. Validate that we should process this repository
		shouldProcess, err := validateRepository(repositoryPath)

		if shouldProcess {
			// 3. Fix permissions
			err = os.Chmod(repositoryPath, os.ModePerm)

			if err != nil {
				os.RemoveAll(repositoryPath)
				writeResponse(responseWriter, http.StatusInternalServerError, "Internal Server Error")
				return
			}

			// 4. Build the docker container
			err = docker(
				"build",
				fmt.Sprintf("--tag=%v", pushEvent.Repository.Name),                  // This allows us to find that container for run
				fmt.Sprintf("--file=%v", path.Join(repositoryPath, dockerFilePath)), // Specify the file as we don't use default Dockerfile
				repositoryPath)

			if err != nil {
				os.RemoveAll(repositoryPath)
				writeResponse(responseWriter, http.StatusInternalServerError, "Internal Server Error")
				return
			}

			// 5. Run the docker container
			repositoryFullPath, err := filepath.Abs(repositoryPath)

			if err != nil {
				os.RemoveAll(repositoryPath)
				docker("rm", pushEvent.Repository.Name)
				writeResponse(responseWriter, http.StatusInternalServerError, "Internal Server Error")
				return
			}

			// docker run --name=test-repository-for-build-my-commit --volume=/home/gildas/projects/go/src/github.com/marmelab/build-my-commit/src/test-repository-for-build-my-commit/:/srv/ test-repository-for-build-my-commit make --file=/srv/Makefile
			err = docker(
				"run",
				fmt.Sprintf("--name=%v", pushEvent.Repository.Name),      // Uses the tag name we specified on build
				fmt.Sprintf("--volume=\"%v:/srv/\"", repositoryFullPath), // Mount the repository inside the container
				pushEvent.Repository.Name,
				"make",
				//"--file=Makefile",										// Specify the Makefile
				"build")

			if err != nil {
				os.RemoveAll(repositoryPath)
				docker("rm", pushEvent.Repository.Name)
				writeResponse(responseWriter, http.StatusInternalServerError, "Internal Server Error")
				return
			}

			// 6. Check wether the ouput has changed with git diff
			output, err := gitWithContext(
				"status",
				repositoryPath,
				"--porcelain") // This will make git status return a machine readable output without pretty formatting

			if err != nil {
				writeResponse(responseWriter, http.StatusInternalServerError, "Internal Server Error")
				os.RemoveAll(repositoryPath)
				docker("rm", pushEvent.Repository.Name)
				return
			}

			// 7. Commit & push the output if necessary
			if output != "" && len(output) > 0 {
				log.Println("Output has changed")
				// Add files as it may be the first build
				_, err = gitWithContext(
					"add",
					repositoryPath,
					"build")

				if err != nil {
					writeResponse(responseWriter, http.StatusInternalServerError, "Internal Server Error")
					os.RemoveAll(repositoryPath)
					docker("rm", pushEvent.Repository.Name)
					return
				}

				// Commit files
				message, err := getCommitMessage(pushEvent.HeadCommit.ID)

				if err != nil {
					writeResponse(responseWriter, http.StatusInternalServerError, "Internal Server Error")
					os.RemoveAll(repositoryPath)
					docker("rm", pushEvent.Repository.Name)
					return
				}

				_, err = gitWithContext(
					"commit",
					repositoryPath,
					fmt.Sprintf("--message=%v", message),
					".")

				if err != nil {
					writeResponse(responseWriter, http.StatusInternalServerError, "Internal Server Error")
					os.RemoveAll(repositoryPath)
					docker("rm", pushEvent.Repository.Name)
					return
				}

				// Ensure the remote repository  hasn't been updated while we built this commit
				lastHash, err := gitWithContext(
					"rev-parse",
					repositoryPath,
					"origin/master")

				if err != nil {
					writeResponse(responseWriter, http.StatusInternalServerError, "Internal Server Error")
					os.RemoveAll(repositoryPath)
					docker("rm", pushEvent.Repository.Name)
					return
				}
				log.Printf("Last commit hash: \"%q\"", lastHash)
				log.Printf("Pushevent commit hash: \"%q\"", pushEvent.HeadCommit.ID)

				// Repository hasn't been updated if we get the same hash as pushEvent
				// If not, we just discard this build as another one is probably running in parallel
				if lastHash == pushEvent.HeadCommit.ID {
					log.Println("Last commit hash hasn't changed")
					// Push files
					_, err = gitWithContext(
						"push",
						repositoryPath)

					if err != nil {
						writeResponse(responseWriter, http.StatusInternalServerError, "Internal Server Error")
						os.RemoveAll(repositoryPath)
						docker("rm", pushEvent.Repository.Name)
						return
					}
				}
			}

			os.RemoveAll(repositoryPath)
			docker("rm", pushEvent.Repository.Name)
			writeResponse(responseWriter, http.StatusOK, "Repository processed")
		} else {
			os.RemoveAll(repositoryPath)
			writeResponse(responseWriter, http.StatusBadRequest, "Cannot process this repository")
			return
		}

		return
	}

	// Request should not be handled, just return
	return
}
