package buildmycommit

import (
	//"bytes"
	//"net/http"
	//"net/http/httptest"
	"github.com/marmelab/buildmycommit/states"
	"testing"
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

	newMachine = func(state int) HookRequestMachine {
		machine := TestMachine{
			StartState:    state,
			StateHandlers: map[int]StateHandler{},
			EndStates:     map[int]bool{}}

		return machine
	}

	HandleHookRequest(nil, nil)
}
