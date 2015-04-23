package main

import (
	"net/http"
)

// Handler define the signature of finite state machine state handlers
type Handler func(State) (int, State, error)

// Machine is the finite state machine (fsm)
type Machine struct {
	Handlers   map[int]Handler
	StartState int
	EndStates  map[int]bool
}

// AddState adds a state and its handler to the fsm
func (machine *Machine) AddState(state int, handlerFn Handler) {
	machine.Handlers[state] = handlerFn
}

// AddEndState adds an end state
func (machine *Machine) AddEndState(endState int) {
	machine.EndStates[endState] = true
}

// Execute the appropriate handler
func (machine *Machine) Execute(state State) {
	if handler, present := machine.Handlers[machine.StartState]; present {
		for {
			nextState, nextStatePayload, err := handler(state)

			// Handle errors globally
			if err != nil {
				state.Status = http.StatusInternalServerError
				state.StatusMessage = "Internal Server Error"
				nextState = StateEndRequest
			}

			_, finished := machine.EndStates[nextState]
			if finished {
				break
			} else {
				handler, present = machine.Handlers[nextState]
				state = nextStatePayload
			}
		}
	}
}
