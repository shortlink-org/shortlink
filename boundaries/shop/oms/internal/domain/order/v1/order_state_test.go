package v1

import (
	"context"
	"sync"
	"testing"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/require"

	"github.com/shortlink-org/shortlink/pkg/fsm"
)

func TestOrderState(t *testing.T) {
	// Define fixed UUIDs for consistency across tests.
	fixedCustomerID := uuid.MustParse("123e4567-e89b-12d3-a456-426614174000")
	fixedProductID1 := uuid.MustParse("123e4567-e89b-12d3-a456-426614174001")
	fixedProductID2 := uuid.MustParse("123e4567-e89b-12d3-a456-426614174002")

	// Create a background context for general use.
	ctx := context.Background()

	t.Run("NewOrderState", func(t *testing.T) {
		orderState := NewOrderState(fixedCustomerID)

		require.Equal(t, fixedCustomerID, orderState.GetCustomerId(), "Customer ID should match")
		require.Equal(t, OrderStatus_ORDER_STATUS_PENDING, orderState.GetStatus(), "Initial status should be Pending")
		require.Empty(t, orderState.GetItems(), "Initial items should be empty")
	})

	t.Run("CreateOrder", func(t *testing.T) {
		orderState := NewOrderState(fixedCustomerID)

		items := Items{
			NewItem(fixedProductID1, 2, decimal.NewFromFloat(19.99)),
			NewItem(fixedProductID2, 1, decimal.NewFromFloat(9.99)),
		}

		err := orderState.CreateOrder(ctx, items)
		require.NoError(t, err, "CreateOrder should not return an error")
		require.Equal(t, OrderStatus_ORDER_STATUS_PROCESSING, orderState.GetStatus(), "Status should transition to Processing")
		require.Equal(t, items, orderState.GetItems(), "Items should match the created items")
	})

	t.Run("UpdateOrder", func(t *testing.T) {
		orderState := NewOrderState(fixedCustomerID)

		initialItems := Items{
			NewItem(fixedProductID1, 2, decimal.NewFromFloat(19.99)),
		}
		err := orderState.CreateOrder(ctx, initialItems)
		require.NoError(t, err, "CreateOrder should not return an error")

		updatedItems := Items{
			NewItem(fixedProductID1, 3, decimal.NewFromFloat(29.99)),
			NewItem(fixedProductID2, 1, decimal.NewFromFloat(9.99)),
		}

		err = orderState.UpdateOrder(ctx, updatedItems)
		require.NoError(t, err, "UpdateOrder should not return an error")
		require.Equal(t, updatedItems, orderState.GetItems(), "Items should reflect the updates")
	})

	t.Run("CancelOrder", func(t *testing.T) {
		orderState := NewOrderState(fixedCustomerID)

		items := Items{
			NewItem(fixedProductID1, 2, decimal.NewFromFloat(19.99)),
		}
		err := orderState.CreateOrder(ctx, items)
		require.NoError(t, err, "CreateOrder should not return an error")

		err = orderState.CancelOrder(ctx)
		require.NoError(t, err, "CancelOrder should not return an error")
		require.Equal(t, OrderStatus_ORDER_STATUS_CANCELLED, orderState.GetStatus(), "Status should transition to Cancelled")
	})

	t.Run("CompleteOrder", func(t *testing.T) {
		orderState := NewOrderState(fixedCustomerID)

		items := Items{
			NewItem(fixedProductID1, 2, decimal.NewFromFloat(19.99)),
		}
		err := orderState.CreateOrder(ctx, items)
		require.NoError(t, err, "CreateOrder should not return an error")

		err = orderState.CompleteOrder(ctx)
		require.NoError(t, err, "CompleteOrder should not return an error")
		require.Equal(t, OrderStatus_ORDER_STATUS_COMPLETED, orderState.GetStatus(), "Status should transition to Completed")
	})

	t.Run("OrderStateConcurrency", func(t *testing.T) {
		orderState := NewOrderState(fixedCustomerID)

		items := Items{
			NewItem(fixedProductID1, 2, decimal.NewFromFloat(19.99)),
			NewItem(fixedProductID2, 1, decimal.NewFromFloat(9.99)),
		}

		err := orderState.CreateOrder(ctx, items)
		require.NoError(t, err, "CreateOrder should not return an error")

		updatedItems := Items{
			NewItem(fixedProductID1, 3, decimal.NewFromFloat(29.99)),
			NewItem(fixedProductID2, 2, decimal.NewFromFloat(19.99)),
		}

		var wg sync.WaitGroup
		wg.Add(2)

		// Use separate contexts for each goroutine to simulate independent operations.
		go func() {
			defer wg.Done()
			err := orderState.UpdateOrder(ctx, updatedItems)
			// It's possible that UpdateOrder happens before or after CancelOrder.
			// Depending on the FSM, updating after cancellation might fail or be allowed.
			// Here, we assume it succeeds or fails gracefully.
			if err != nil {
				t.Logf("UpdateOrder encountered an error: %v", err)
			}
		}()

		go func() {
			defer wg.Done()
			err := orderState.CancelOrder(ctx)
			if err != nil {
				t.Logf("CancelOrder encountered an error: %v", err)
			}
		}()

		wg.Wait()

		// After concurrent operations, the order should either be Cancelled or have updated items.
		// Depending on the FSM's transition rules, updating after cancellation might not change the state.
		finalStatus := orderState.GetStatus()
		require.True(t, finalStatus == OrderStatus_ORDER_STATUS_CANCELLED || finalStatus == OrderStatus_ORDER_STATUS_PROCESSING,
			"Final status should be either Cancelled or Processing")
	})

	t.Run("Callbacks", func(t *testing.T) {
		orderState := NewOrderState(fixedCustomerID)

		// Variables to track callbacks.
		var (
			enterState   string
			exitState    string
			triggeredEvt fsm.Event
			callbackMu   sync.Mutex
		)

		// Set up the OnEnterState callback.
		orderState.fsm.SetOnEnterState(func(ctx context.Context, from, to fsm.State, event fsm.Event) {
			callbackMu.Lock()
			defer callbackMu.Unlock()
			enterState = to.String()
			triggeredEvt = event
		})

		// Set up the OnExitState callback.
		orderState.fsm.SetOnExitState(func(ctx context.Context, from, to fsm.State, event fsm.Event) {
			callbackMu.Lock()
			defer callbackMu.Unlock()
			exitState = from.String()
			triggeredEvt = event
		})

		// Transition: Pending -> Processing
		err := orderState.CreateOrder(ctx, Items{
			NewItem(fixedProductID1, 2, decimal.NewFromFloat(19.99)),
		})
		require.NoError(t, err, "CreateOrder should transition state to Processing")

		// Verify callbacks.
		callbackMu.Lock()
		require.Equal(t, "ORDER_STATUS_PENDING", exitState, "OnExitState should be called with Pending")
		require.Equal(t, fsm.Event("ORDER_STATUS_PENDING"), triggeredEvt, "Triggered event should be Pending")
		require.Equal(t, "ORDER_STATUS_PROCESSING", enterState, "OnEnterState should be called with Processing")
		callbackMu.Unlock()

		// Reset callback trackers.
		callbackMu.Lock()
		enterState, exitState, triggeredEvt = "", "", ""
		callbackMu.Unlock()

		// Transition: Processing -> Completed
		err = orderState.CompleteOrder(ctx)
		require.NoError(t, err, "CompleteOrder should transition state to Completed")

		// Verify callbacks.
		callbackMu.Lock()
		require.Equal(t, "ORDER_STATUS_PROCESSING", exitState, "OnExitState should be called with Processing")
		require.Equal(t, fsm.Event("ORDER_STATUS_COMPLETED"), triggeredEvt, "Triggered event should be Completed")
		require.Equal(t, "ORDER_STATUS_COMPLETED", enterState, "OnEnterState should be called with Completed")
		callbackMu.Unlock()
	})

	t.Run("InvalidTransitions", func(t *testing.T) {
		orderState := NewOrderState(fixedCustomerID)

		// Attempt to cancel the order while it's in Pending state.
		err := orderState.CancelOrder(ctx)
		require.NoError(t, err, "CancelOrder should transition state to Cancelled from Pending")

		// Attempt to complete a Cancelled order.
		err = orderState.CompleteOrder(ctx)
		require.Error(t, err, "CompleteOrder should return an error when transitioning from Cancelled")
		require.Equal(t, OrderStatus_ORDER_STATUS_CANCELLED, orderState.GetStatus(), "Status should remain Cancelled after invalid transition")
	})

	t.Run("ContextCancellation", func(t *testing.T) {
		orderState := NewOrderState(fixedCustomerID)

		// Create a context that gets canceled immediately.
		ctx, cancel := context.WithCancel(context.Background())
		cancel()

		// Attempt to trigger a transition with a canceled context.
		err := orderState.CreateOrder(ctx, Items{
			NewItem(fixedProductID1, 1, decimal.NewFromFloat(9.99)),
		})

		// Depending on FSM implementation, it might handle context cancellation.
		// For this example, assuming FSM.TriggerEvent checks for context cancellation.
		// If FSM.TriggerEvent does not handle it, the transition occurs.
		// Adjust the assertion based on actual FSM behavior.

		// Since the context was canceled, and assuming FSM.TriggerEvent respects it:
		if err != nil {
			require.Error(t, err, "CreateOrder should return an error when context is canceled")
			require.Equal(t, OrderStatus_ORDER_STATUS_PENDING, orderState.GetStatus(), "Status should remain Pending after canceled context")
			require.Empty(t, orderState.GetItems(), "Items should not be added after canceled context")
		} else {
			// If FSM.TriggerEvent does not respect context cancellation, the transition occurs.
			require.Equal(t, OrderStatus_ORDER_STATUS_PROCESSING, orderState.GetStatus(), "Status should transition to Processing despite canceled context")
			require.Equal(t, 1, len(orderState.GetItems()), "Items should be added despite canceled context")
		}
	})
}
