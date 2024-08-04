package v1

import (
	"context"
	"sync"
	"testing"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func TestOrderState(t *testing.T) {
	ctx := context.Background()

	fixedCustomerID := uuid.MustParse("123e4567-e89b-12d3-a456-426614174000")
	fixedProductID1 := uuid.MustParse("123e4567-e89b-12d3-a456-426614174001")
	fixedProductID2 := uuid.MustParse("123e4567-e89b-12d3-a456-426614174002")

	t.Run("NewOrderState", func(t *testing.T) {
		orderState := NewOrderState(fixedCustomerID)

		assert.Equal(t, fixedCustomerID, orderState.GetCustomerId())
		assert.Equal(t, OrderStatus_ORDER_STATUS_PENDING, orderState.GetStatus())
		assert.Empty(t, orderState.GetItems())
	})

	t.Run("CreateOrder", func(t *testing.T) {
		orderState := NewOrderState(fixedCustomerID)

		items := Items{
			NewItem(fixedProductID1, 2, decimal.NewFromFloat(19.99)),
			NewItem(fixedProductID2, 1, decimal.NewFromFloat(9.99)),
		}

		err := orderState.CreateOrder(ctx, items)
		assert.NoError(t, err)
		assert.Equal(t, OrderStatus_ORDER_STATUS_PROCESSING, orderState.GetStatus())
		assert.Equal(t, items, orderState.GetItems())
	})

	t.Run("UpdateOrder", func(t *testing.T) {
		orderState := NewOrderState(fixedCustomerID)

		initialItems := Items{
			NewItem(fixedProductID1, 2, decimal.NewFromFloat(19.99)),
		}
		err := orderState.CreateOrder(ctx, initialItems)
		assert.NoError(t, err)

		updatedItems := Items{
			NewItem(fixedProductID1, 3, decimal.NewFromFloat(29.99)),
			NewItem(fixedProductID2, 1, decimal.NewFromFloat(9.99)),
		}

		err = orderState.UpdateOrder(ctx, updatedItems)
		assert.NoError(t, err)
		assert.Equal(t, updatedItems, orderState.GetItems())
	})

	t.Run("CancelOrder", func(t *testing.T) {
		orderState := NewOrderState(fixedCustomerID)

		items := Items{
			NewItem(fixedProductID1, 2, decimal.NewFromFloat(19.99)),
		}
		err := orderState.CreateOrder(ctx, items)
		assert.NoError(t, err)

		err = orderState.CancelOrder(ctx)
		assert.NoError(t, err)
		assert.Equal(t, OrderStatus_ORDER_STATUS_CANCELLED, orderState.GetStatus())
	})

	t.Run("CompleteOrder", func(t *testing.T) {
		orderState := NewOrderState(fixedCustomerID)

		items := Items{
			NewItem(fixedProductID1, 2, decimal.NewFromFloat(19.99)),
		}
		err := orderState.CreateOrder(ctx, items)
		assert.NoError(t, err)

		err = orderState.CompleteOrder(ctx)
		assert.NoError(t, err)
		assert.Equal(t, OrderStatus_ORDER_STATUS_COMPLETED, orderState.GetStatus())
	})

	t.Run("OrderStateConcurrency", func(t *testing.T) {
		orderState := NewOrderState(fixedCustomerID)

		items := Items{
			NewItem(fixedProductID1, 2, decimal.NewFromFloat(19.99)),
			NewItem(fixedProductID2, 1, decimal.NewFromFloat(9.99)),
		}

		err := orderState.CreateOrder(ctx, items)
		assert.NoError(t, err)

		updatedItems := Items{
			NewItem(fixedProductID1, 3, decimal.NewFromFloat(29.99)),
			NewItem(fixedProductID2, 1, decimal.NewFromFloat(9.99)),
		}

		var wg sync.WaitGroup
		wg.Add(2)

		go func() {
			defer wg.Done()
			_ = orderState.UpdateOrder(ctx, updatedItems)
		}()

		go func() {
			defer wg.Done()
			_ = orderState.CancelOrder(ctx)
		}()

		wg.Wait()

		assert.True(t, orderState.GetStatus() == OrderStatus_ORDER_STATUS_CANCELLED ||
			orderState.GetStatus() == OrderStatus_ORDER_STATUS_PROCESSING)
	})
}
