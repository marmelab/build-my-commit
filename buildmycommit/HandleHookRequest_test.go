package buildmycommit

import (
	"reflect"
	"testing"

	"github.com/marmelab/buildmycommit/states"
)

func TestNewStateReturnsAStateWithARandomlyGeneratedRepositoryPath(t *testing.T) {
	state := newState(nil, nil)

	if state.RepositoryPath == "" {
		t.Errorf("new state should have a valid repository path")
	}
}

// TestMachine is the finite state machine (fsm)
type TestMachine struct {
	StateHandlers map[int]StateHandler
	StartState    int
	EndStates     map[int]bool
}

// AddState adds a state and its handler to the fsm
func (machine TestMachine) AddState(state int, stateHandler StateHandler) {
	machine.StateHandlers[state] = stateHandler
}

// AddEndState adds an end state
func (machine TestMachine) AddEndState(endState int) {
	machine.EndStates[endState] = true

}

// Execute the appropriate handler
func (machine TestMachine) Execute(state states.State) error {
	return nil
}

func TestHandleHookRequestShouldInitializeANewMachineWithStateValidateRequest(t *testing.T) {

	newMachine = func(state int) HookRequestMachine {
		expectedState := states.ValidateRequest

		if state != expectedState {
			t.Errorf("initial state should be %v", expectedState)
		}

		machine := TestMachine{
			StartState:    state,
			StateHandlers: map[int]StateHandler{},
			EndStates:     map[int]bool{}}

		return machine
	}

	HandleHookRequest(nil, nil)
}

func TestHandleHookRequestShouldAddStatesHandlersOnCorrectStates(t *testing.T) {
	var testMachine TestMachine
	var errorMessage = "expected %v as handler for state %v, found %v"

	newMachine = func(state int) HookRequestMachine {
		machine := TestMachine{
			StartState:    state,
			StateHandlers: map[int]StateHandler{},
			EndStates:     map[int]bool{}}

		testMachine = machine
		return machine
	}

	HandleHookRequest(nil, nil)

	if len(testMachine.StateHandlers) != 15 {
		t.Errorf("testMachine should have %v state handlers", 15)
	}

	handler := reflect.TypeOf(testMachine.StateHandlers[states.ValidateRequest]).Elem().Name()
	if handler != "ValidateRequest" {
		t.Errorf(errorMessage, "ValidateRequest", states.ValidateRequest, handler)
	}

	handler = reflect.TypeOf(testMachine.StateHandlers[states.ParsePayload]).Elem().Name()
	if handler != "ParsePayload" {
		t.Errorf(errorMessage, "ParsePayload", states.ParsePayload, handler)
	}

	handler = reflect.TypeOf(testMachine.StateHandlers[states.ValidatePushEvent]).Elem().Name()
	if handler != "ValidatePushEvent" {
		t.Errorf(errorMessage, "ValidatePushEvent", states.ValidatePushEvent, handler)
	}

	handler = reflect.TypeOf(testMachine.StateHandlers[states.CloneRepository]).Elem().Name()
	if handler != "CloneRepository" {
		t.Errorf(errorMessage, "CloneRepository", states.CloneRepository, handler)
	}

	handler = reflect.TypeOf(testMachine.StateHandlers[states.ValidateDockerFile]).Elem().Name()
	if handler != "ValidateDockerFile" {
		t.Errorf(errorMessage, "ValidateDockerFile", states.ValidateDockerFile, handler)
	}

	handler = reflect.TypeOf(testMachine.StateHandlers[states.ValidateCommit]).Elem().Name()
	if handler != "ValidateCommit" {
		t.Errorf(errorMessage, "ValidateCommit", states.ValidateCommit, handler)
	}

	handler = reflect.TypeOf(testMachine.StateHandlers[states.BuildDocker]).Elem().Name()
	if handler != "BuildDocker" {
		t.Errorf(errorMessage, "BuildDocker", states.BuildDocker, handler)
	}

	handler = reflect.TypeOf(testMachine.StateHandlers[states.RunDocker]).Elem().Name()
	if handler != "RunDocker" {
		t.Errorf(errorMessage, "RunDocker", states.RunDocker, handler)
	}

	handler = reflect.TypeOf(testMachine.StateHandlers[states.CompareBuild]).Elem().Name()
	if handler != "CompareBuild" {
		t.Errorf(errorMessage, "CompareBuild", states.CompareBuild, handler)
	}

	handler = reflect.TypeOf(testMachine.StateHandlers[states.CommitBuild]).Elem().Name()
	if handler != "CommitBuild" {
		t.Errorf(errorMessage, "CommitBuild", states.CommitBuild, handler)
	}

	handler = reflect.TypeOf(testMachine.StateHandlers[states.CheckRemoteHash]).Elem().Name()
	if handler != "CheckRemoteHash" {
		t.Errorf(errorMessage, "CheckRemoteHash", states.CheckRemoteHash, handler)
	}

	handler = reflect.TypeOf(testMachine.StateHandlers[states.PushBuild]).Elem().Name()
	if handler != "PushBuild" {
		t.Errorf(errorMessage, "PushBuild", states.PushBuild, handler)
	}

	handler = reflect.TypeOf(testMachine.StateHandlers[states.EndRequest]).Elem().Name()
	if handler != "EndRequest" {
		t.Errorf(errorMessage, "EndRequest", states.EndRequest, handler)
	}

	handler = reflect.TypeOf(testMachine.StateHandlers[states.CleanRepository]).Elem().Name()
	if handler != "CleanRepository" {
		t.Errorf(errorMessage, "CleanRepository", states.CleanRepository, handler)
	}

	handler = reflect.TypeOf(testMachine.StateHandlers[states.CleanDocker]).Elem().Name()
	if handler != "CleanDocker" {
		t.Errorf(errorMessage, "CleanDocker", states.CleanDocker, handler)
	}
}

func TestHandleHookRequestShouldAddStateRequestHandledAsEndState(t *testing.T) {
	var testMachine TestMachine

	newMachine = func(state int) HookRequestMachine {
		machine := TestMachine{
			StartState:    state,
			StateHandlers: map[int]StateHandler{},
			EndStates:     map[int]bool{}}

		testMachine = machine
		return machine
	}

	HandleHookRequest(nil, nil)

	if !testMachine.EndStates[states.RequestHandled] {
		t.Errorf("End states should be %v", states.RequestHandled)
	}
}
