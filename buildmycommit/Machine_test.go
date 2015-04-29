package buildmycommit

import (
	"testing"

	"github.com/marmelab/buildmycommit/states"
)

type TestHandlerForEndStateTesting struct {
}

func (stateHandler TestHandlerForEndStateTesting) Handle(state states.State) (int, states.State) {
	return 42, states.State{}
}

type TestHandler struct {
	t *testing.T
}

func (stateHandler TestHandler) Handle(state states.State) (int, states.State) {
	return 1, states.State{Status: 24}
}

type TestHandler2 struct {
	t *testing.T
}

func (stateHandler TestHandler2) Handle(state states.State) (int, states.State) {
	if state.Status != 24 {
		stateHandler.t.Errorf("Execute should pass states between handlers")
	}

	return 42, states.State{Status: 42}
}

func TestNewMachineReturnsANewMachineWithSpecifiedStartState(t *testing.T) {
	var machine Machine
	machine = NewMachine(42).(Machine)

	if machine.StartState != 42 {
		t.Errorf("NewMachine should return a new machine initialized with StartState at 42")
	}
}

func TestNewMachineReturnsANewMachineWithStateHandlersMap(t *testing.T) {
	var machine Machine
	machine = NewMachine(42).(Machine)

	if machine.StateHandlers == nil {
		t.Errorf("NewMachine should return a new machine initialized with StateHandlers map")
	}
}

func TestNewMachineReturnsANewMachineWithEndStatesMap(t *testing.T) {
	var machine Machine
	machine = NewMachine(42).(Machine)

	if machine.EndStates == nil {
		t.Errorf("NewMachine should return a new machine initialized with EndStates map")
	}
}

func TestAddStateAddsTheState(t *testing.T) {
	machine := Machine{
		StartState:    42,
		StateHandlers: map[int]StateHandler{},
		EndStates:     map[int]bool{}}

	machine.AddState(0, TestHandler{})

	if len(machine.StateHandlers) != 1 {
		t.Errorf("AddState should have added the state handler to the StateHandlers map")
	}
}

func TestAddEndStateAddsTheState(t *testing.T) {
	machine := Machine{
		StartState:    42,
		StateHandlers: map[int]StateHandler{},
		EndStates:     map[int]bool{}}

	machine.AddEndState(42)

	if len(machine.EndStates) != 1 {
		t.Errorf("AddEndState should have added the state to the EndStates map")
	}
}

func TestExecuteShouldReturnWhenReachingEndState(t *testing.T) {
	machine := Machine{
		StartState:    0,
		StateHandlers: map[int]StateHandler{},
		EndStates:     map[int]bool{}}

	machine.AddState(0, TestHandlerForEndStateTesting{})
	machine.AddEndState(42)

	// This test will timeout if failing
	machine.Execute(states.State{})
}

func TestExecuteShouldPassStatesBetweenHandlers(t *testing.T) {
	machine := Machine{
		StartState:    0,
		StateHandlers: map[int]StateHandler{},
		EndStates:     map[int]bool{}}

	machine.AddState(0, TestHandler{t: t})
	machine.AddState(1, TestHandler2{t: t})
	machine.AddEndState(42)

	// This test will timeout if failing
	machine.Execute(states.State{})
}

func TestExecuteShouldFailIfAnInvalidStateIsReached(t *testing.T) {
	machine := Machine{
		StartState:    0,
		StateHandlers: map[int]StateHandler{},
		EndStates:     map[int]bool{}}

	machine.AddState(0, TestHandler{t: t})

	// This test will timeout if failing
	err := machine.Execute(states.State{})

	if err == nil {
		t.Errorf("Execute should have returned an error")
	}
}
