package buildmycommit

import (
	"github.com/marmelab/buildmycommit/states"
	"log"
	"os"
)

// StateHandler define the interface for state handlers
type StateHandler interface {
	Handle(states.State) (int, states.State)
}

// Machine is the finite state machine (fsm)
type Machine struct {
	StateHandlers map[int]StateHandler
	StartState    int
	EndStates     map[int]bool
}

// AddState adds a state and its handler to the fsm
func (machine *Machine) AddState(state int, stateHandler StateHandler) {
	machine.StateHandlers[state] = stateHandler
}

// AddEndState adds an end state
func (machine *Machine) AddEndState(endState int) {
	machine.EndStates[endState] = true
}

// Execute the appropriate handler
func (machine *Machine) Execute(state states.State) {
	if handler, present := machine.StateHandlers[machine.StartState]; present {
		for {
			nextState, nextStatePayload := handler.Handle(state)

			debug := os.Getenv("DEBUG")
			if len(debug) > 0 {
				log.Printf("[MACHINE]: state=%v payload=%v", nextState, state)
			}

			_, finished := machine.EndStates[nextState]
			if finished {
				break
			} else {
				handler, present = machine.StateHandlers[nextState]
				state = nextStatePayload
			}
		}
	}
}
