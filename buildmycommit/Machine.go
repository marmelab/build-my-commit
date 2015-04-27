package buildmycommit

import (
	"fmt"
	"github.com/marmelab/buildmycommit/states"
)

// StateHandler define the interface for state handlers
type StateHandler interface {
	Handle(states.State) (int, states.State)
}

// HookRequestMachine is the interface for a Finite State Machine handling hook requests processing
type HookRequestMachine interface {
	AddState(state int, stateHandler StateHandler)
	AddEndState(endState int)
	Execute(state states.State) error
}

// Machine is the finite state machine (fsm)
type Machine struct {
	StateHandlers map[int]StateHandler
	StartState    int
	EndStates     map[int]bool
}

// AddState adds a state and its handler to the fsm
func (machine Machine) AddState(state int, stateHandler StateHandler) {
	machine.StateHandlers[state] = stateHandler
}

// AddEndState adds an end state
func (machine Machine) AddEndState(endState int) {
	machine.EndStates[endState] = true

}

// Execute the appropriate handler
func (machine Machine) Execute(state states.State) error {
	if handler, present := machine.StateHandlers[machine.StartState]; present {
		for {

			nextState, nextStatePayload := handler.Handle(state)

			_, finished := machine.EndStates[nextState]
			if finished {
				break
			} else {
				handler, present = machine.StateHandlers[nextState]
				if handler == nil {
					return fmt.Errorf("No handler found for state %v", machine.StartState)
				}
				state = nextStatePayload
			}
		}
	}

	return nil
}

// NewMachine returns a new Machine instance
func NewMachine(initialState int) HookRequestMachine {
	machine := Machine{
		StartState:    initialState,
		StateHandlers: map[int]StateHandler{},
		EndStates:     map[int]bool{}}

	return machine
}
