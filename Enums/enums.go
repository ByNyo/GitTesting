package main

import "fmt"

type ServerState int // Creates a new Type called ServerState

const ( // Defines the types of ServerStates
	StateIdle = iota
	StateConnected
	StateError
	StateRetrying
)

var stateName = map[ServerState]string{ // Map variable stateName
	StateIdle:      "idle", // Sets the state name of stateidle
	StateConnected: "connected",
	StateError:     "error",
	StateRetrying:  "retrying",
}

func (ss ServerState) String() string { // Function that takes a ServerState as parameter
	return stateName[ss] // Returns the StateName as string
}

func main() {
	ns := transition(StateIdle) // Transitions the state to StateIdle and saves it in a variable
	fmt.Println(ns)             // Prints the current ServerState

	ns2 := transition(ns) // Transitions the state to the state of the variabnle ns and saves it in a variable
	fmt.Println(ns2)      // Prints the current ServerState
}

func transition(s ServerState) ServerState {
	switch s { // Switch for ServerState s
	case StateIdle: // Checks if ServerState s is Stateidle and executes case if its
		return StateConnected // Returns a ServerState
	case StateConnected, StateRetrying: // Checks if ServerState s is StateConnected or StateRetrying and executes case if its
		return StateIdle // Returns a ServerState
	case StateError: // Checks if ServerState s is StateError and executes case if its
		return StateError // Returns a ServerState
	default:
		panic(fmt.Errorf("unknown state: %s", s)) // Panics showing an error with an/the unknown state
	}
}
