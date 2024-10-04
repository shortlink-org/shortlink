package fsm

import (
	"fmt"
	"sync"
)

// State represents a state in the FSM.
type State string

// Event represents an event that can trigger a state transition.
type Event string

// TransitionRuleSet defines the allowed state transitions.
// It maps a State to a map of Events and their resulting States.
type TransitionRuleSet map[State]map[Event]State

// FSM represents a finite state machine.
type FSM struct {
	mu sync.RWMutex

	CurrentState    State
	TransitionRules TransitionRuleSet

	// Callbacks triggered on state transitions.
	OnEnterState func(from State, to State, event Event)
	OnExitState  func(from State, to State, event Event)
}

// New creates a new FSM with the given initial state.
// It initializes the TransitionRules to prevent nil map assignments.
func New(initialState State) *FSM {
	return &FSM{
		CurrentState:    initialState,
		TransitionRules: make(TransitionRuleSet),
	}
}

// AddTransitionRule adds a transition rule to the FSM.
// It defines that when in the 'from' state, upon receiving 'event',
// the FSM should transition to the 'to' state.
func (f *FSM) AddTransitionRule(from State, event Event, to State) {
	f.mu.Lock()
	defer f.mu.Unlock()

	if _, exists := f.TransitionRules[from]; !exists {
		f.TransitionRules[from] = make(map[Event]State)
	}

	f.TransitionRules[from][event] = to
}

// SetOnEnterState sets the callback function to be called when entering a new state.
func (f *FSM) SetOnEnterState(callback func(from State, to State, event Event)) {
	f.mu.Lock()
	defer f.mu.Unlock()
	f.OnEnterState = callback
}

// SetOnExitState sets the callback function to be called when exiting a state.
func (f *FSM) SetOnExitState(callback func(from State, to State, event Event)) {
	f.mu.Lock()
	defer f.mu.Unlock()
	f.OnExitState = callback
}

// TriggerEvent triggers an event and attempts to transition the FSM to the next state.
// It returns an error if the transition is invalid.
func (f *FSM) TriggerEvent(event Event) error {
	f.mu.Lock()
	defer f.mu.Unlock()

	current := f.CurrentState

	// Retrieve the next state based on the current state and event.
	nextState, valid := f.TransitionRules[current][event]
	if !valid {
		return fmt.Errorf("invalid transition: no rule for event '%s' in state '%s'", event, current)
	}

	// If there is an exit callback, invoke it before changing the state.
	if f.OnExitState != nil {
		f.OnExitState(current, nextState, event)
	}

	// Update the current state.
	f.CurrentState = nextState

	// If there is an enter callback, invoke it after changing the state.
	if f.OnEnterState != nil {
		f.OnEnterState(current, nextState, event)
	}

	return nil
}

// GetCurrentState returns the current state of the FSM.
// It uses a read lock to allow concurrent reads.
func (f *FSM) GetCurrentState() State {
	f.mu.RLock()
	defer f.mu.RUnlock()
	return f.CurrentState
}
