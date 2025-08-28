//go:build unit

package saga

import (
	"context"
	"fmt"
	"os"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/goleak"

	"github.com/shortlink-org/go-sdk/logger"
)

type Wallet struct {
	mu sync.Mutex

	value int
}

func TestMain(m *testing.M) {
	goleak.VerifyTestMain(m)

	os.Exit(m.Run())
}

func TestNewSaga(t *testing.T) {
	// Init logger
	conf := config.Configuration{
		Level: config.DEBUG_LEVEL,
	}
	log, err := logger.New(logger.Zap, conf)
	require.NoError(t, err, "Error init a logger")

	t.Run("create simple saga", func(t *testing.T) {
		const SAGA_NAME = "Number magic"
		const SAGA_STEP_A = "A"
		const SAGA_STEP_B = "B"
		const SAGA_STEP_C = "C"
		const SAGA_STEP_D = "D"
		const SAGA_STEP_E = "E"

		ctx := context.Background()

		// Example amount
		wallet := &Wallet{
			value: 0,
		}

		addFunc := func(ctx context.Context) error {
			wallet.mu.Lock()
			wallet.value += 10
			wallet.mu.Unlock()
			return nil
		}
		rejectFunc := func(ctx context.Context, thenError error) error {
			wallet.mu.Lock()
			wallet.value -= 10
			wallet.mu.Unlock()
			return nil
		}
		printFunc := func(ctx context.Context) error {
			log.Info(fmt.Sprintf("amount: %d", wallet.value))
			return nil
		}
		rejectPrintFunc := func(ctx context.Context, thenError error) error {
			log.Info(fmt.Sprintf("amount: %d", wallet.value))
			return nil
		}

		// create a new saga for work with number
		sagaNumber, errs := New(SAGA_NAME, SetLogger(log)).
			WithContext(ctx).
			Build()
		// check error
		assert.Len(t, errs, 0)
		if len(errs) > 0 {
			t.Fatal(errs)
		}

		// add step increment
		_, errs = sagaNumber.AddStep(SAGA_STEP_A).
			Then(addFunc).
			Reject(rejectFunc).
			Build()
		// check error
		assert.Len(t, errs, 0)
		if len(errs) > 0 {
			t.Fatal(errs)
		}

		// add step decrement
		_, errs = sagaNumber.AddStep(SAGA_STEP_B).
			Then(addFunc).
			Reject(rejectFunc).
			Build()
		// check error
		assert.Len(t, errs, 0)
		if len(errs) > 0 {
			t.Fatal(errs)
		}

		// add step decrement
		_, errs = sagaNumber.AddStep(SAGA_STEP_C).
			Then(printFunc).
			Reject(rejectPrintFunc).
			Needs(SAGA_STEP_A, SAGA_STEP_B).
			Build()
		// check error
		assert.Len(t, errs, 0)
		if len(errs) > 0 {
			t.Fatal(errs)
		}

		// add step decrement
		_, errs = sagaNumber.AddStep(SAGA_STEP_D).
			Then(addFunc).
			Reject(rejectFunc).
			Needs(SAGA_STEP_C).
			Build()
		// check error
		assert.Len(t, errs, 0)
		if len(errs) > 0 {
			t.Fatal(errs)
		}

		// add step decrement
		_, errs = sagaNumber.AddStep(SAGA_STEP_E).
			Then(printFunc).
			Reject(rejectPrintFunc).
			Needs(SAGA_STEP_D, SAGA_STEP_B).
			Build()
		// check error
		assert.Len(t, errs, 0)
		if len(errs) > 0 {
			t.Fatal(errs)
		}

		// Run saga
		err := sagaNumber.Play(nil)
		assert.Equal(t, wallet.value, 30)
		require.NoError(t, err)
	})

	t.Run("create simple saga with reject", func(t *testing.T) {
		const SAGA_NAME = "Number magic"
		const SAGA_STEP_A = "A"
		const SAGA_STEP_B = "B"
		const SAGA_STEP_C = "C"
		const SAGA_STEP_D = "D"
		const SAGA_STEP_E = "E"
		const SAGA_STEP_FAIL = "STEP_FAIL"

		ctx := context.Background()

		// Example amount
		wallet := &Wallet{
			value: 0,
		}

		addFunc := func(ctx context.Context) error {
			wallet.mu.Lock()
			wallet.value += 10
			wallet.mu.Unlock()
			return nil
		}
		rejectFunc := func(ctx context.Context, therErr error) error {
			wallet.mu.Lock()
			wallet.value -= 9 // For check work addFunc after saga.Play ;-)
			wallet.mu.Unlock()
			return nil
		}
		printFunc := func(ctx context.Context) error {
			log.Info(fmt.Sprintf("amount: %d", wallet.value))
			return nil
		}
		rejectPrintFunc := func(ctx context.Context, thenError error) error {
			log.Info(fmt.Sprintf("amount: %d", wallet.value))
			return nil
		}

		// create a new saga for work with number
		sagaNumber, errs := New(SAGA_NAME, SetLogger(log)).
			WithContext(ctx).
			Build()
		// check error
		assert.Len(t, errs, 0)
		if len(errs) > 0 {
			t.Fatal(errs)
		}

		// add step increment
		_, errs = sagaNumber.AddStep(SAGA_STEP_A).
			Then(addFunc).
			Reject(rejectFunc).
			Build()
		// check error
		assert.Len(t, errs, 0)
		if len(errs) > 0 {
			t.Fatal(errs)
		}

		// add step decrement
		_, errs = sagaNumber.AddStep(SAGA_STEP_B).
			Then(addFunc).
			Reject(rejectFunc).
			Build()
		// check error
		assert.Len(t, errs, 0)
		if len(errs) > 0 {
			t.Fatal(errs)
		}

		// add step decrement
		_, errs = sagaNumber.AddStep(SAGA_STEP_C).
			Then(printFunc).
			Reject(rejectPrintFunc).
			Needs(SAGA_STEP_A, SAGA_STEP_B).
			Build()
		// check error
		assert.Len(t, errs, 0)
		if len(errs) > 0 {
			t.Fatal(errs)
		}

		// add step decrement
		_, errs = sagaNumber.AddStep(SAGA_STEP_D).
			Then(addFunc).
			Reject(rejectFunc).
			Needs(SAGA_STEP_C).
			Build()
		// check error
		assert.Len(t, errs, 0)
		if len(errs) > 0 {
			t.Fatal(errs)
		}

		// add step decrement
		_, errs = sagaNumber.AddStep(SAGA_STEP_E).
			Then(printFunc).
			Reject(rejectPrintFunc).
			Needs(SAGA_STEP_D, SAGA_STEP_B).
			Build()
		// check error
		assert.Len(t, errs, 0)
		if len(errs) > 0 {
			t.Fatal(errs)
		}

		// add step decrement
		_, errs = sagaNumber.AddStep(SAGA_STEP_FAIL).
			Then(func(ctx context.Context) error {
				return fmt.Errorf("SAGA_STEP_FAIL")
			}).
			Reject(rejectPrintFunc).
			Needs(SAGA_STEP_E, SAGA_STEP_C).
			Build()
		// check error
		assert.Len(t, errs, 0)
		if len(errs) > 0 {
			t.Fatal(errs)
		}

		// Run saga
		err := sagaNumber.Play(nil)
		assert.Equal(t, wallet.value, 3) // amount: 10+10+10-9-9-9=3
		require.NoError(t, err)
	})

	t.Run("check error case", func(t *testing.T) {
		const SAGA_NAME = "Error Check Saga"
		const SAGA_STEP_A = "A"
		const SAGA_STEP_B = "B"

		ctx := context.Background()

		wallet := &Wallet{
			value: 0,
		}

		// Functions to simulate success and failure
		successFunc := func(ctx context.Context) error {
			wallet.mu.Lock()
			wallet.value += 10
			wallet.mu.Unlock()
			return nil
		}
		rejectSuccessFunc := func(ctx context.Context, thenErr error) error {
			wallet.mu.Lock()
			wallet.value -= 10
			wallet.mu.Unlock()
			return nil
		}
		failFunc := func(ctx context.Context) error {
			return fmt.Errorf("forced error")
		}

		// Create a new saga
		saga, errs := New(SAGA_NAME, SetLogger(log)).
			WithContext(ctx).
			Build()
		assert.Len(t, errs, 0)

		// Add step A
		_, errs = saga.AddStep(SAGA_STEP_A).
			Then(successFunc).
			Reject(rejectSuccessFunc).
			Build()
		assert.Len(t, errs, 0)

		// Add step B, which will force an error
		_, errs = saga.AddStep(SAGA_STEP_B).
			Then(failFunc).
			Reject(func(ctx context.Context, thenErr error) error {
				return thenErr
			}).
			Build()
		assert.Len(t, errs, 0)

		// Run saga
		err := saga.Play(nil)
		require.Error(t, err)                           // Ensure that an error is returned
		assert.Contains(t, err.Error(), "forced error") // Check if the error message is as expected
		assert.Equal(t, wallet.value, 0)                // The wallet value should remain 0 as the saga should rollback
	})
}
