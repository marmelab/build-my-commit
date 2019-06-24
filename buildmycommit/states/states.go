package states

const (
	// ValidateRequest : State called when a new request is received
	ValidateRequest = iota
	// ParsePayload : State called to parse the json payload
	ParsePayload = iota
	// ValidatePushEvent : State called when a new push event payload is received
	ValidatePushEvent = iota
	// CloneRepository : State called to clone the git repository
	CloneRepository = iota
	// ValidateDockerFile : State called to validate that the repository has a build.Dockerfile
	ValidateDockerFile = iota
	// ValidateCommit : State called to validate that the commit trigering the build is not from an automatic build
	ValidateCommit = iota
	// BuildDocker : State called to build the docker container
	BuildDocker = iota
	// RunDocker : State called to run the docker container
	RunDocker = iota
	// CompareBuild : State called to compare built output with remote build
	CompareBuild = iota
	// CommitBuild : State called to commit the built output
	CommitBuild = iota
	// CheckRemoteHash : State called to verify that the remote branch has not been modified since the build
	CheckRemoteHash = iota
	// PushBuild : State called to push the built output
	PushBuild = iota
	// CleanRepository : State called to clean the cloned repository by deleting the local directory
	CleanRepository = iota
	// CleanDocker : State called to clean the docker container by remving it
	CleanDocker = iota
	// EndRequest : State called to end request
	EndRequest = iota
	// RequestHandled : State called when processing has been completed
	RequestHandled = iota
	// InternalServerError : State called when an internal server error occured
	InternalServerError = iota
)
