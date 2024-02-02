package fsm

import (
	"testing"
)

func TestFSM(t *testing.T) {
	const (
		state1 State = "state1"
		state2 State = "state2"
		state3 State = "state3"
	)

	// Initialize a new FSM
	fsm := NewFSM("state1")

	// Add some transition rules
	fsm.AddTransitionRule(state1, "event1", state1)
	fsm.AddTransitionRule(state2, "event2", state3)

	// Variables to track callbacks
	var entered, exited State

	// Define callbacks
	fsm.OnEnterState = func(state State) {
		entered = state
	}
	fsm.OnExitState = func(state State) {
		exited = state
	}

	// Trigger the first event
	fsm.TriggerEvent("event1")

	// Test the state transition and callbacks
	if fsm.CurrentState != state2 {
		t.Errorf("Expected state2, got %s", fsm.CurrentState)
	}
	if entered != "state2" {
		t.Errorf("Expected entered state2, got %s", entered)
	}
	if exited != "state1" {
		t.Errorf("Expected exited state1, got %s", exited)
	}

	// Reset callback trackers
	entered, exited = "", ""

	// Trigger the second event
	fsm.TriggerEvent("event2")

	// Test the state transition and callbacks
	if fsm.CurrentState != state3 {
		t.Errorf("Expected state3, got %s", fsm.CurrentState)
	}
	if entered != "state3" {
		t.Errorf("Expected entered state3, got %s", entered)
	}
	if exited != "state2" {
		t.Errorf("Expected exited state2, got %s", exited)
	}

	// Test invalid event
	fsm.TriggerEvent("invalid")

	// State should not change on invalid event
	if fsm.CurrentState != state3 {
		t.Errorf("Expected state3, got %s", fsm.CurrentState)
	}
}
