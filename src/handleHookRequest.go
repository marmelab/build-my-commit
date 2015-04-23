package main

import (
	"net/http"
)

// DockerFilePath is the default Dockerfile to use for automatic build
const DockerFilePath = "build.Dockerfile"

// CommitMessage is the default commit message used to push builded output
const CommitMessage = "[build-my-commmit] Updated build from"

const (
	// StateRequestReceived : State called when a new request is received
	StateRequestReceived = iota
	// StateParsePayload : State called to parse the json payload
	StateParsePayload = iota
	// StatePushEventReceived : State called when a new push event payload is received
	StatePushEventReceived = iota
	// StateCloneRepository : State called to clone the git repository
	StateCloneRepository = iota
	// StateValidateRepository : State called to validate that the repository should be processed
	StateValidateRepository = iota
	// StateBuildDocker : State called to build the docker container
	StateBuildDocker = iota
	// StateRunDocker : State called to run the docker container
	StateRunDocker = iota
	// StateCompareOutput : State called to compare built output with remote build
	StateCompareOutput = iota
	// StateCommitOutput : State called to commit the built output
	StateCommitOutput = iota
	// StateVerifyRemoteHash : State called to verify that the remote branch has not been modified since the build
	StateVerifyRemoteHash = iota
	// StatePushOutput : State called to push the built output
	StatePushOutput = iota
	// StateCleanGit : State called to clean the cloned repository by deleting the local directory
	StateCleanGit = iota
	// StateCleanDocker : State called to clean the docker container by remving it
	StateCleanDocker = iota
	// StateEndRequest : State called to end request
	StateEndRequest = iota
	// StateRequestHandled : State called when processing has been completed
	StateRequestHandled = iota
)

// State holds the data to pass between each states
type State struct {
	ResponseWriter    http.ResponseWriter
	Request           *http.Request
	PushEvent         PushEvent
	RepositoryPath    string
	Status            int
	StatusMessage     string
	ShouldCleanGit    bool
	ShouldCleanDocker bool
}

func handleHookRequest(responseWriter http.ResponseWriter, request *http.Request) {
	state := State{ResponseWriter: responseWriter, Request: request}

	machine := Machine{map[int]Handler{}, StatePushEventReceived, map[int]bool{}}

	machine.AddState(StateRequestReceived, HandleStateRequestReceived)
	machine.AddState(StateParsePayload, HandleStateParsePayload)
	machine.AddState(StatePushEventReceived, HandleStatePushEventReceived)
	machine.AddState(StateCloneRepository, HandleStateCloneRepository)
	machine.AddState(StateValidateRepository, HandleStateValidateRepository)
	machine.AddState(StateBuildDocker, HandleStateBuildDocker)
	machine.AddState(StateRunDocker, HandleStateRunDocker)
	machine.AddState(StateCompareOutput, HandleStateCompareOutput)
	machine.AddState(StateCommitOutput, HandleStateCommitOutput)
	machine.AddState(StateVerifyRemoteHash, HandleStateVerifyRemoteHash)
	machine.AddState(StatePushOutput, HandleStatePushOutput)
	machine.AddState(StateEndRequest, HandleStateEndRequest)
	machine.AddState(StateCleanGit, HandleStateCleanGit)
	machine.AddState(StateCleanDocker, HandleStateCleanDocker)
	machine.AddEndState(StateRequestHandled)

	machine.Execute(state)
	return
}
