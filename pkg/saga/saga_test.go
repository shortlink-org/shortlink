// +build unit

package saga

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/goleak"

	"github.com/batazor/shortlink/pkg/saga/store/ram"
)

func TestMain(m *testing.M) {
	goleak.VerifyTestMain(m)
}

func TestNewSaga(t *testing.T) {
	t.Run("create simple saga", func(t *testing.T) {
		const SAGA_NAME = "Number magic"
		const SAGA_STEP_A = "A"
		const SAGA_STEP_B = "B"
		const SAGA_STEP_C = "C"
		const SAGA_STEP_D = "D"
		const SAGA_STEP_E = "E"

		ctx := context.Background()
		store := ram.RAM{}

		// Example amount
		amount := 0

		addFunc := func(ctx context.Context) error {
			amount += 10
			return nil
		}
		rejectFunc := func(ctx context.Context) error {
			amount -= 10
			return nil
		}
		printFunc := func(ctx context.Context) error {
			fmt.Println(amount)
			return nil
		}

		// create a new saga for work with number
		sagaNumber, errs := New(SAGA_NAME).
			WithContext(ctx).
			SetStore(store).
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
		assert.Equal(t, amount, 30)

		assert.Nil(t, err)
	})
}
