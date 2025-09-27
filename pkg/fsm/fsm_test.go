package fsm

import (
	"context"
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
)

// TestTrafficLightFSM verifies the functionality of the Traffic Light FSM with context support.
func TestTrafficLightFSM(t *testing.T) {
	t.Attr("type", "unit")
	t.Attr("package", "fsm")
	t.Attr("component", "fsm")

	// Define traffic light states.
	const (
		StateRed    State = "Red"
		StateGreen  State = "Green"
		StateYellow State = "Yellow"
	)

	// Define traffic light events.
	const (
		EventTimer     Event = "Timer"
		EventEmergency Event = "Emergency" // Invalid event for this FSM
	)

	// Initialize the FSM with the initial state (Red).
	trafficLight := New(StateRed)

	// Add transition rules to the FSM.
	trafficLight.AddTransitionRule(StateRed, EventTimer, StateGreen)
	trafficLight.AddTransitionRule(StateGreen, EventTimer, StateYellow)
	trafficLight.AddTransitionRule(StateYellow, EventTimer, StateRed)

	// Variables to track callback invocations.
	var (
		entered      State
		exited       State
		triggeredEv  Event
		callbackLock sync.Mutex
	)

	// Set up the OnExitState callback.
	trafficLight.SetOnExitState(func(ctx context.Context, from, to State, event Event) {
		callbackLock.Lock()
		defer callbackLock.Unlock()
		exited = from
		triggeredEv = event
	})

	// Set up the OnEnterState callback.
	trafficLight.SetOnEnterState(func(ctx context.Context, from, to State, event Event) {
		callbackLock.Lock()
		defer callbackLock.Unlock()
		entered = to
		triggeredEv = event
	})

	// Helper function to reset callback trackers.
	resetCallbacks := func() {
		callbackLock.Lock()
		defer callbackLock.Unlock()
		entered, exited, triggeredEv = "", "", ""
	}

	// Create a context for the test.
	ctx := t.Context()

	// Test initial state.
	require.Equal(t, StateRed, trafficLight.GetCurrentState(), "Initial state should be Red")

	// Transition 1: Red -> Green
	err := trafficLight.TriggerEvent(ctx, EventTimer)
	require.NoError(t, err, "TriggerEvent should not return an error for valid event 'Timer' from Red")

	// Verify state transition to Green.
	require.Equal(t, StateGreen, trafficLight.GetCurrentState(), "State should transition to Green after Timer event")

	// Verify callbacks for the first transition.
	require.Equal(t, StateRed, exited, "Exited state should be Red after Timer event")
	require.Equal(t, StateGreen, entered, "Entered state should be Green after Timer event")
	require.Equal(t, EventTimer, triggeredEv, "Triggered event should be Timer")

	// Reset callback trackers.
	resetCallbacks()

	// Transition 2: Green -> Yellow
	err = trafficLight.TriggerEvent(ctx, EventTimer)
	require.NoError(t, err, "TriggerEvent should not return an error for valid event 'Timer' from Green")

	// Verify state transition to Yellow.
	require.Equal(t, StateYellow, trafficLight.GetCurrentState(), "State should transition to Yellow after Timer event")

	// Verify callbacks for the second transition.
	require.Equal(t, StateGreen, exited, "Exited state should be Green after Timer event")
	require.Equal(t, StateYellow, entered, "Entered state should be Yellow after Timer event")
	require.Equal(t, EventTimer, triggeredEv, "Triggered event should be Timer")

	// Reset callback trackers.
	resetCallbacks()

	// Transition 3: Yellow -> Red
	err = trafficLight.TriggerEvent(ctx, EventTimer)
	require.NoError(t, err, "TriggerEvent should not return an error for valid event 'Timer' from Yellow")

	// Verify state transition to Red.
	require.Equal(t, StateRed, trafficLight.GetCurrentState(), "State should transition to Red after Timer event")

	// Verify callbacks for the third transition.
	require.Equal(t, StateYellow, exited, "Exited state should be Yellow after Timer event")
	require.Equal(t, StateRed, entered, "Entered state should be Red after Timer event")
	require.Equal(t, EventTimer, triggeredEv, "Triggered event should be Timer")

	// Reset callback trackers.
	resetCallbacks()

	// Attempt to trigger an invalid event: Emergency
	err = trafficLight.TriggerEvent(ctx, EventEmergency)
	require.Error(t, err, "TriggerEvent should return an error for invalid event 'Emergency'")
	require.Contains(t, err.Error(), "invalid transition", "Error message should indicate invalid transition")

	// Verify that the state remains unchanged after invalid event.
	require.Equal(t, StateRed, trafficLight.GetCurrentState(), "State should remain Red after invalid Emergency event")

	// Verify that callbacks were not called on invalid event.
	require.Equal(t, State(""), exited, "Exited state should not be set on invalid Emergency event")
	require.Equal(t, State(""), entered, "Entered state should not be set on invalid Emergency event")
	require.Equal(t, Event(""), triggeredEv, "Triggered event should not be set on invalid Emergency event")

	// Ensure FSM is still operational by triggering another valid event: Timer
	err = trafficLight.TriggerEvent(ctx, EventTimer)
	require.NoError(t, err, "TriggerEvent should not return an error for valid event 'Timer' after invalid event")

	// Verify state transition to Green.
	require.Equal(t, StateGreen, trafficLight.GetCurrentState(), "State should transition to Green after Timer event")

	// Verify callbacks for the transition after invalid event.
	require.Equal(t, StateRed, exited, "Exited state should be Red after Timer event")
	require.Equal(t, StateGreen, entered, "Entered state should be Green after Timer event")
	require.Equal(t, EventTimer, triggeredEv, "Triggered event should be Timer")
}
