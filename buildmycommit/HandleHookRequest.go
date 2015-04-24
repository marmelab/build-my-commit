package buildmycommit

import (
	"encoding/json"
	"github.com/marmelab/buildmycommit/docker"
	"github.com/marmelab/buildmycommit/git"
	"github.com/marmelab/buildmycommit/statehandlers"
	"github.com/marmelab/buildmycommit/states"
	"github.com/marmelab/buildmycommit/tools"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

const dockerFilePath = "build.Dockerfile"
const commitMessage = "[build-my-commmit] Updated build from"

// HandleHookRequest handles the git webhooks requests
func HandleHookRequest(responseWriter http.ResponseWriter, request *http.Request) {
	state := states.State{
		ResponseWriter: responseWriter,
		Request:        request,
		RepositoryPath: tools.GenerateRandomID()}

	machine := Machine{
		StateHandlers: map[int]StateHandler{},
		StartState:    states.ValidateRequest,
		EndStates:     map[int]bool{}}

	machine.AddState(states.ValidateRequest, statehandlers.NewValidateRequest("Invalid HTTP method"))
	machine.AddState(states.ParsePayload, statehandlers.NewParsePayload(ioutil.ReadAll, json.Unmarshal, "Invalid webhook payload"))
	machine.AddState(states.ValidatePushEvent, statehandlers.NewValidatePushEvent("refs/heads/master", "Discarded: not on master"))
	machine.AddState(states.CloneRepository, statehandlers.NewCloneRepository(git.Git))
	machine.AddState(states.ValidateDockerFile, statehandlers.NewValidateDockerFile(tools.Exists, dockerFilePath, "Discarded: no build.Dockerfile for automatic build"))
	machine.AddState(states.ValidateCommit, statehandlers.NewValidateCommit(git.GitWithContext, commitMessage, "Discarded: detected commit by automatic build"))
	machine.AddState(states.BuildDocker, statehandlers.NewBuildDocker(docker.Docker, dockerFilePath))
	machine.AddState(states.RunDocker, statehandlers.NewRunDocker(docker.Docker, filepath.Abs))
	machine.AddState(states.CompareBuild, statehandlers.NewCompareBuild(git.GitWithContext, "Discarded: build already up to date on master"))
	machine.AddState(states.CommitBuild, statehandlers.NewCommitBuild(git.GitWithContext, commitMessage))
	machine.AddState(states.CheckRemoteHash, statehandlers.NewCheckRemoteHash(git.GitWithContext, "Discarded: master has been updated"))
	machine.AddState(states.PushBuild, statehandlers.NewPushBuild(git.GitWithContext))
	machine.AddState(states.EndRequest, statehandlers.NewEndRequest())
	machine.AddState(states.CleanRepository, statehandlers.NewCleanRepository(os.RemoveAll))
	machine.AddState(states.CleanDocker, statehandlers.NewCleanDocker(docker.Docker))
	machine.AddEndState(states.RequestHandled)

	machine.Execute(state)
	return
}
