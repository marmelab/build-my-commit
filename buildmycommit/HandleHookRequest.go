package buildmycommit

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"

	"github.com/marmelab/buildmycommit/docker"
	"github.com/marmelab/buildmycommit/git"
	"github.com/marmelab/buildmycommit/statehandlers"
	"github.com/marmelab/buildmycommit/states"
	"github.com/marmelab/buildmycommit/tools"
)

const dockerFilePath = "build.Dockerfile"
const commitMessage = "[build-my-commmit] Updated build from"

func newState(responseWriter http.ResponseWriter, request *http.Request) states.State {
	state := states.State{
		ResponseWriter: responseWriter,
		Request:        request,
		RepositoryPath: tools.GenerateRandomID()}

	return state
}

var newMachine = NewMachine

// HandleHookRequest handles the git webhooks requests
func HandleHookRequest(responseWriter http.ResponseWriter, request *http.Request) {
	state := newState(responseWriter, request)

	dockerExec := docker.GetDockerCmd()
	gitExec := git.GetGitCmd()
	machine := newMachine(states.ValidateRequest)

	machine.AddState(states.ValidateRequest, statehandlers.NewValidateRequest("Invalid HTTP method"))
	machine.AddState(states.ParsePayload, statehandlers.NewParsePayload(ioutil.ReadAll, json.Unmarshal, "Invalid webhook payload"))
	machine.AddState(states.ValidatePushEvent, statehandlers.NewValidatePushEvent("refs/heads/master", "Discarded: not on master"))
	machine.AddState(states.CloneRepository, statehandlers.NewCloneRepository(gitExec.Exec))
	machine.AddState(states.ValidateDockerFile, statehandlers.NewValidateDockerFile(tools.Exists, dockerFilePath, "Discarded: no build.Dockerfile for automatic build"))
	machine.AddState(states.ValidateCommit, statehandlers.NewValidateCommit(gitExec.ExecInContext, commitMessage, "Discarded: detected commit by automatic build"))
	machine.AddState(states.BuildDocker, statehandlers.NewBuildDocker(dockerExec.Exec, dockerFilePath))
	machine.AddState(states.RunDocker, statehandlers.NewRunDocker(dockerExec.Exec, filepath.Abs))
	machine.AddState(states.CompareBuild, statehandlers.NewCompareBuild(gitExec.ExecInContext, "Discarded: build already up to date on master"))
	machine.AddState(states.CommitBuild, statehandlers.NewCommitBuild(gitExec.ExecInContext, commitMessage))
	machine.AddState(states.CheckRemoteHash, statehandlers.NewCheckRemoteHash(gitExec.ExecInContext, "Discarded: master has been updated"))
	machine.AddState(states.PushBuild, statehandlers.NewPushBuild(gitExec.ExecInContext))
	machine.AddState(states.EndRequest, statehandlers.NewEndRequest())
	machine.AddState(states.CleanRepository, statehandlers.NewCleanRepository(os.RemoveAll))
	machine.AddState(states.CleanDocker, statehandlers.NewCleanDocker(dockerExec.Exec))
	machine.AddEndState(states.RequestHandled)

	machine.Execute(state)
	return
}
