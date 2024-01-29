package fsm

import (
	"sync"
)

// State type for representing states in the FSM
type State string

// Event type for representing events that trigger state transitions
type Event string

// TransitionRuleSet defines the allowed state transitions
type TransitionRuleSet map[State]map[Event]State

// FSM represents a finite state machine
type FSM struct {
	mu sync.Mutex

	CurrentState    State
	TransitionRules TransitionRuleSet

	// Callbacks
	OnEnterState func(state State)
	OnExitState  func(state State)
}

// TriggerEvent triggers an event and changes the state of the FSM accordingly
func (f *FSM) TriggerEvent(event Event) {
	f.mu.Lock()
	defer f.mu.Unlock()

	nextState, ok := f.TransitionRules[f.CurrentState][event]
	if !ok {
		// No transition rule for this event in the current state
		return
	}

	// Exit the current state
	if f.OnExitState != nil {
		f.OnExitState(f.CurrentState)
	}

	// Update the current state
	f.CurrentState = nextState

	// Enter the new state
	if f.OnEnterState != nil {
		f.OnEnterState(f.CurrentState)
	}
}

// AddTransitionRule adds a transition rule to the FSM
func (f *FSM) AddTransitionRule(from State, event Event, to State) {
	f.mu.Lock()
	defer f.mu.Unlock()

	if f.TransitionRules == nil {
		f.TransitionRules = make(TransitionRuleSet)
	}

	if _, ok := f.TransitionRules[from]; !ok {
		f.TransitionRules[from] = make(map[Event]State)
	}

	f.TransitionRules[from][event] = to
}

// NewFSM creates a new FSM with the given initial state
func NewFSM(initialState State) *FSM {
	return &FSM{
		CurrentState:    initialState,
		TransitionRules: make(TransitionRuleSet),
	}
}
