//go:build unit
// +build unit

package saga

import (
	"context"
	"fmt"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/goleak"

	"github.com/batazor/shortlink/internal/pkg/logger"
)

type Wallet struct {
	sync.Mutex

	value int
}

func TestMain(m *testing.M) {
	goleak.VerifyTestMain(m)
}

func TestNewSaga(t *testing.T) {
	// Init logger
	conf := logger.Configuration{
		Level: logger.DEBUG_LEVEL,
	}
	log, err := logger.NewLogger(logger.Zap, conf)
	assert.Nil(t, err, "Error init a logger")

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
			wallet.Lock()
			wallet.value += 10
			wallet.Unlock()
			return nil
		}
		rejectFunc := func(ctx context.Context) error {
			wallet.Lock()
			wallet.value -= 10
			wallet.Unlock()
			return nil
		}
		printFunc := func(ctx context.Context) error {
			log.Info(fmt.Sprintf("amount: %d", wallet.value))
			return nil
		}

		// create a new saga for work with number
		sagaNumber, errs := New(SAGA_NAME, Logger(log)).
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
			Reject(printFunc).
			Needs([]string{SAGA_STEP_A, SAGA_STEP_B}).
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
			Needs([]string{SAGA_STEP_C}).
			Build()
		// check error
		assert.Len(t, errs, 0)
		if len(errs) > 0 {
			t.Fatal(errs)
		}

		// add step decrement
		_, errs = sagaNumber.AddStep(SAGA_STEP_E).
			Then(printFunc).
			Reject(printFunc).
			Needs([]string{SAGA_STEP_D, SAGA_STEP_B}).
			Build()
		// check error
		assert.Len(t, errs, 0)
		if len(errs) > 0 {
			t.Fatal(errs)
		}

		// Run saga
		err := sagaNumber.Play(nil)
		assert.Equal(t, wallet.value, 30)
		assert.Nil(t, err)
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
			wallet.Lock()
			wallet.value += 10
			wallet.Unlock()
			return nil
		}
		rejectFunc := func(ctx context.Context) error {
			wallet.Lock()
			wallet.value -= 9 // For check work addFunc after saga.Play ;-)
			wallet.Unlock()
			return nil
		}
		printFunc := func(ctx context.Context) error {
			log.Info(fmt.Sprintf("amount: %d", wallet.value))
			return nil
		}

		// create a new saga for work with number
		sagaNumber, errs := New(SAGA_NAME, Logger(log)).
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
			Reject(printFunc).
			Needs([]string{SAGA_STEP_A, SAGA_STEP_B}).
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
			Needs([]string{SAGA_STEP_C}).
			Build()
		// check error
		assert.Len(t, errs, 0)
		if len(errs) > 0 {
			t.Fatal(errs)
		}

		// add step decrement
		_, errs = sagaNumber.AddStep(SAGA_STEP_E).
			Then(printFunc).
			Reject(printFunc).
			Needs([]string{SAGA_STEP_D, SAGA_STEP_B}).
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
			Reject(printFunc).
			Needs([]string{SAGA_STEP_E, SAGA_STEP_C}).
			Build()
		// check error
		assert.Len(t, errs, 0)
		if len(errs) > 0 {
			t.Fatal(errs)
		}

		// Run saga
		err := sagaNumber.Play(nil)
		assert.Equal(t, wallet.value, 3) // amount: 10+10+10-9-9-9=3
		assert.Nil(t, err)
	})
}
